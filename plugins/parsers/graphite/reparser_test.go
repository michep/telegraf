package graphite

import (
	"fmt"
	"testing"
)

var (
	re_templates = []string{
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>Gate)\\.(?P<gatecomponent>(?:ifm\\.\\w+?)|(?:\\w+?))\\.(?P<name>\\w+?Transmitter)-(?P<peer>\\w+?\\d+?)\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>Gate)\\.(?P<gatecomponent>(?:ifm\\.\\w+?)|(?:\\w+?))\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<name>Connector)(?:\\d*?)(?P<_name>DatabaseAccessor)\\.(?:\\w*?)(?P<name>(?:SmsPostMessage|SmsGetMessage).+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<name>Connector)(?:\\d*?)(?P<_name>DatabaseAccessor)\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>RcoiConnectorInMessageTransmitter)\\.(?P<peer>\\w+?\\d+?)\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>\\w+?Transmitter)-(?P<peer>\\w+?\\d+?)\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>\\w+?Receiver)-(?P<peer>\\w+?\\d+?)\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>\\w+?Adapter)-(?P<type>[\\w]+?)\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>.+?commandStatusMonitorAvgThroughputCounter)\\.(?P<status>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>.+?channelInMessageProcessMonitorAvgThroughputCounter)\\.(?P<subject>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>.+?PercentileCounter)\\.(?P<percentile>\\w+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>.+?)\\.(?:priority)\\.(?P<priority>\\w+?)\\.(?P<name>.+$)",
		"(?P<name>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>\\w+?)\\.(?P<name>.+$)",
	}
)

func BenchmarkParseReParser(b *testing.B) {
	p, _ := NewGraphiteReParser(".", re_templates, nil)

	for i := 0; i < b.N; i++ {
		p.ApplyTemplate("MfmsMonitor.manager-base-sbmanager3.zsbmng03.UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.priority.6.size")
	}
}

func TestTemplateApplyReParser(t *testing.T) {
	var tests = []struct {
		input       string
		template    string
		measurement string
		tags        map[string]string
		err         string
	}{
		{
			input:       "MfmsMonitor.avirouter-base-avirouter0.zchl09.AdvisaOutMessageCache.advisaOutMessageCachedMap.size",
			measurement: "MfmsMonitor.AdvisaOutMessageCache.advisaOutMessageCachedMap.size",
			tags:        map[string]string{"component": "avirouter-base-avirouter0", "zone": "zchl09"},
		},
		{
			input:       "MfmsMonitor.imsichannel-megalabs-megalabs0.zchl04.ProtocolAdapter-agroros.undefinedMsisdnResponseMonitorAvgThroughputCounter",
			measurement: "MfmsMonitor.ProtocolAdapter.undefinedMsisdnResponseMonitorAvgThroughputCounter",
			tags:        map[string]string{"component": "imsichannel-megalabs-megalabs0", "zone": "zchl04", "type": "agroros"},
		},
		{
			input:       "MfmsMonitor.manager-base-sbmanager3.zsbmng03.ComiConnectorOutMessageReceiver-sb8.receivedConnectorOutMessageMonitorAvgThroughputCounter",
			measurement: "MfmsMonitor.ComiConnectorOutMessageReceiver.receivedConnectorOutMessageMonitorAvgThroughputCounter",
			tags:        map[string]string{"component": "manager-base-sbmanager3", "zone": "zsbmng03", "peer": "sb8"},
		},
		{
			input:       "MfmsMonitor.channel-smpp-alfacapmts0.zchl10.CmiChannelStateTransmitter-sbmanager1.channelStateProcessQueueProcessor.size",
			measurement: "MfmsMonitor.CmiChannelStateTransmitter.channelStateProcessQueueProcessor.size",
			tags:        map[string]string{"component": "channel-smpp-alfacapmts0", "zone": "zchl10", "peer": "sbmanager1"},
		},
		{
			input:       "MfmsMonitor.connector-emailfileex-vtb2414.zcnr08.EmailFileExConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter.io",
			measurement: "MfmsMonitor.ConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter.io",
			tags:        map[string]string{"component": "connector-emailfileex-vtb2414", "zone": "zcnr08"},
		},
		{
			input:       "MfmsMonitor.connector-alfa5-alfa15.zcnr00.AlfaConnector5DatabaseAccessor.connectorImsiResponseProcStatusAddQueueProcessor.size",
			measurement: "MfmsMonitor.ConnectorDatabaseAccessor.connectorImsiResponseProcStatusAddQueueProcessor.size",
			tags:        map[string]string{"component": "connector-alfa5-alfa15", "zone": "zcnr00"},
		},
		{
			input:       "MfmsMonitor.connector-fileex-russta3.zcnr00.FileExConnectorDatabaseAccessor.fileExSmsPostMessageAddQueueProcessor.size",
			measurement: "MfmsMonitor.ConnectorDatabaseAccessor.SmsPostMessageAddQueueProcessor.size",
			tags:        map[string]string{"component": "connector-fileex-russta3", "zone": "zcnr00"},
		},
		{
			input:       "MfmsMonitor.connector-hpx-autoins1.zcnr08.HpxConnector0DatabaseAccessor.hpxSmsGetMessageAddMonitorAvgSpeedCounter",
			measurement: "MfmsMonitor.ConnectorDatabaseAccessor.SmsGetMessageAddMonitorAvgSpeedCounter",
			tags:        map[string]string{"component": "connector-hpx-autoins1", "zone": "zcnr08"},
		},
		{
			input:       "MfmsMonitor.connector-sb1-sb13.zcnr02.Gate.infobip0n1.CmiChannelInMessageTransmitterManager.channelInMessageProcessMonitorAvgSpeedCounter",
			measurement: "MfmsMonitor.Gate.CmiChannelInMessageTransmitterManager.channelInMessageProcessMonitorAvgSpeedCounter",
			tags:        map[string]string{"component": "connector-sb1-sb13", "zone": "zcnr02", "gatecomponent": "infobip0n1"},
		},
		{
			input:       "MfmsMonitor.connector-sb1-sb2.zsbcnr01.Gate.ifm.ws0n5.MonitorAccessor.monitorParameterProcessQueueProcessor.size",
			measurement: "MfmsMonitor.Gate.MonitorAccessor.monitorParameterProcessQueueProcessor.size",
			tags:        map[string]string{"component": "connector-sb1-sb2", "zone": "zsbcnr01", "gatecomponent": "ifm.ws0n5"},
		},
		{
			input:       "MfmsMonitor.connector-sb1-sb13.zcnr02.Gate.ermb0n0.ComiConnectorOutMessageTransmitter-manager1n0.connectorOutMessageProcessMonitorAvgSpeedCounter",
			measurement: "MfmsMonitor.Gate.ComiConnectorOutMessageTransmitter.connectorOutMessageProcessMonitorAvgSpeedCounter",
			tags:        map[string]string{"component": "connector-sb1-sb13", "zone": "zcnr02", "gatecomponent": "ermb0n0", "peer": "manager1n0"},
		},
		{
			input:       "MfmsMonitor.receiver-base-receiver1.zchl04.RcoiConnectorInMessageTransmitter.binbank5.connectorInMessageProcessQueueProcessor.size",
			measurement: "MfmsMonitor.RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
			tags:        map[string]string{"component": "receiver-base-receiver1", "zone": "zchl04", "peer": "binbank5"},
		},
		{
			input:       "MfmsMonitor.connector-emp-mospark1.zcnr03.ComiConnectorOutMessageTransmitterManager.processedConnectorOutMessageMonitorPercentileCounter.90",
			measurement: "MfmsMonitor.ComiConnectorOutMessageTransmitterManager.processedConnectorOutMessageMonitorPercentileCounter",
			tags:        map[string]string{"component": "connector-emp-mospark1", "zone": "zcnr03", "percentile": "90"},
		},
		{
			input:       "MfmsMonitor.manager-base-sbmanager3.zsbmng03.UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.priority.6.size",
			measurement: "MfmsMonitor.UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.size",
			tags:        map[string]string{"component": "manager-base-sbmanager3", "zone": "zsbmng03", "priority": "6"},
		},
		{
			input:       "MfmsMonitor.channel-smpp-alfacapmts0.zchl10.ResendProcessor.commandStatusMonitorAvgThroughputCounter.error",
			measurement: "MfmsMonitor.ResendProcessor.commandStatusMonitorAvgThroughputCounter",
			tags:        map[string]string{"component": "channel-smpp-alfacapmts0", "zone": "zchl10", "status": "error"},
		},
		{
			input:       "MfmsMonitor.channel-smpp-beeline1.zchl06.DeliverSmsProcessor.channelInMessageProcessMonitorAvgThroughputCounter.79037676761",
			measurement: "MfmsMonitor.DeliverSmsProcessor.channelInMessageProcessMonitorAvgThroughputCounter",
			tags:        map[string]string{"component": "channel-smpp-beeline1", "zone": "zchl06", "subject": "79037676761"},
		},
		{
			input:       "MfmsMonitor.receiver-base-receiver0.zchl06.RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
			measurement: "MfmsMonitor.RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
			tags:        map[string]string{"component": "receiver-base-receiver0", "zone": "zchl06"},
		},
	}

	p, _ := NewGraphiteReParser(".", re_templates, nil)

	for _, test := range tests {
		measurement, tags, _, _ := p.ApplyTemplate(test.input)
		fmt.Println(measurement, tags)
		if measurement != test.measurement {
			t.Fatalf("name parse failer.  expected %v, got %v", test.measurement, measurement)
		}
		if len(tags) != len(test.tags) {
			t.Fatalf("unexpected number of tags.  expected %v, got %v", test.tags, tags)
		}
		for k, v := range test.tags {
			if tags[k] != v {
				t.Fatalf("unexpected tag value for tags[%s].  expected %q, got %q", k, v, tags[k])
			}
		}
	}
}
