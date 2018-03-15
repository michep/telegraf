package graphite

import "testing"

var (
	re_templates = []string{
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?P<measurement>Gate)\\.(?P<gatecomponent>(?:ifm\\.[\\w]+?)|(?:[\\w]+?))\\.(?P<measurement>[\\w]+?Transmitter)-(?P<peer>[\\w]+?)\\.(?P<measurement>.+$)",
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?P<measurement>Gate)\\.(?P<gatecomponent>(?:ifm\\.[\\w]+?)|(?:[\\w]+?))\\.(?P<measurement>.+$)",
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?:[\\w]*?)(?P<measurement>Connector)(?:\\d*?)(?P<_measurement>DatabaseAccessor)\\.(?:\\w*?)(?P<measurement>(?:SmsPostMessage|SmsGetMessage).+$)",
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?:[\\w]*?)(?P<measurement>Connector)(?:\\d*?)(?P<_measurement>DatabaseAccessor)\\.(?P<measurement>.+$)",
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?P<measurement>[\\w]+?Transmitter)-(?P<peer>[\\w]+?)\\.(?P<measurement>.+$)",
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?P<measurement>[\\w]+?Receiver)-(?P<peer>[\\w]+?)\\.(?P<measurement>.+$)",
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?P<measurement>[\\w]+?Adapter)-(?P<type>[\\w]+?)\\.(?P<measurement>.+$)",
		"(?P<measurement>^MfmsMonitor)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<zone>[\\w-]+?)\\.(?P<measurement>.+$)",
	}
)

func BenchmarkParseReParser(b *testing.B) {
	p, _ := NewGraphiteReParser(".", re_templates, nil)

	for i := 0; i < b.N; i++ {
		p.ApplyTemplate("MfmsMonitor.avirouter-base-avirouter0.zchl09.AdvisaOutMessageCache.advisaOutMessageCachedMap.size")
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
	}

	p, _ := NewGraphiteReParser(".", re_templates, nil)

	for _, test := range tests {
		measurement, tags, _, _ := p.ApplyTemplate(test.input)
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
