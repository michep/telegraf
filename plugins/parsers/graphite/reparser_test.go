package graphite

import (
	"fmt"
	"testing"
)

var (
	re_templates = map[string][]string{
		"sms": {
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>Gate\\.route)\\.(?P<system>[\\w-]+?)\\.(?P<chl_group>.+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<metric2>.+?TimeCounter)\\.(?P<time>\\w+?$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<metric2>.+?[Mm]essageQueue)\\.(?P<queue>\\w+?)\\.(?P<metric3>.+$) => ${metric1}.${metric2}.${metric3}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<metric2>\\w+?Transmitter)-(?P<peer>\\w+?\\d+?)\\.(?P<metric3>.+$) => ${metric1}.${metric2}.${metric3}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<metric2>.+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<metric1>Connector(Batch)?)(?:\\d*?)(?P<metric2>DatabaseAccessor)\\.(?:\\w*?)(?P<metric3>(SmsPost|SmsGet).+$) => ${metric1}${metric2}.${metric3}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<metric1>Connector)(?:\\d*?)(?P<metric2>DatabaseAccessor)\\.(?P<metric3>.+$) => ${metric1}${metric2}.${metric3}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<metric1>(SmsPost|SmsGet)\\w+)\\.(?:\\w*?)(?P<metric2>(SmsPost|SmsGet).+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?[Mm]essageQueue)\\.(?P<queue>\\w+?)\\.(?P<metric2>.+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>RcoiConnectorInMessageTransmitter)\\.(?P<peer>\\w+?\\d+?)\\.(?P<metric2>.+$) => ${metric1}.peer.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>telegramBot)-(?P<bot>[\\w-]+?)\\.(?P<metric2>.+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>ClientProcessor)\\.(?P<peer>[\\w]+?)\\.(?P<metric2>.+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>\\w+?Transmitter)(?:-|\\.)(?P<peer>\\w+?\\d+?)\\.(?P<metric2>.+$) => ${metric1}.peer.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>\\w+?Receiver)-(?P<peer>[\\w-]+?)\\.(?P<metric2>.+$) => ${metric1}.peer.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>\\w+?Adapter)-(?P<type>[\\w]+?)\\.(?P<metric2>.+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?deliveryMonitorDurationCounter)\\.(?P<type>\\w+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?deliveryDurationMonitorCounter)\\.(?P<type>\\w+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?commandStatusMonitorAvgThroughputCounter)\\.(?P<status>.+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?receivedConnectorOutMessageDlvEventMonitorAvgThroughputCounter)\\.(?P<status>.+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?payloadOutPacketQueue)\\.(?P<peer>[\\w]+?)\\.(?P<metric2>.+$) => ${metric1}.peer.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?channelInMessageProcessMonitorAvgThroughputCounter)\\.(?P<subject>.+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?PercentileCounter)\\.(?P<percentile>\\w+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?(lastRequestDate|lastMessageTime|messagesPerMinute))\\.(?P<peer>\\w+$) => ${metric1}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?(operator|source))\\.(?P<peer>[\\w-]+)\\.(?P<metric2>period)\\.(?P<period>\\w+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?(operator|source))\\.(?P<peer>[\\w-]+)\\.(?P<metric2>status)\\.(?P<status>\\w+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?\\.priority)\\.(?P<priority>\\w+?)\\.(?P<metric2>.+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+?)\\.(?:\\.+)(?P<metric2>.+$) => ${metric1}.${metric2}",
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>.+$) => ${metric1}",
		},
		"push": {
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>StoredQueue)\\.(?P<queue>.+?)\\.(?P<metric2>\\w+$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>ClientOutMessageDlvStatusCounter)\\.(?P<status>[\\w]+?)\\.(?P<type>\\w+$) => ${metric1}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>ConnectorOutMessageDlvStatusCounter)\\.(?P<status>[\\w]+?)\\.(?P<type>\\w+$) => ${metric1}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>Transmitter.*?)\\.(?P<peer>([\\w-]+?(\\d+|demo))|sb)\\.(?P<metric2>(size|queueSize|consumerAvailable)$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>Transmitter.*?)\\.(?P<peer>([\\w-]+?(\\d+|demo))|sb)\\.(?P<metric2>.+?)\\.(?P<type>\\w+$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>In_message.+)\\.cnrId-(?P<cnrid>\\d+?)\\.acnId-(?P<acnid>\\d+)\\.(?P<type>\\w+$) => ${metric1}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+Sender)\\.(?P<peer>[\\w-]+?\\d+)\\.(?P<metric2>.+?)\\.(?P<type>\\w+$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+Sender)\\.(?P<peer>[\\w-]+?\\d+)\\.(?P<metric2>\\w+$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+Processor)-(?P<peer>[\\w\\d-]+?)\\.(?P<metric2>(size|queueSize)$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+Processor)-(?P<peer>[\\w\\d-]+?)\\.(?P<metric2>.+?)\\.(?P<type>\\w+$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+?)\\.acnId-(?P<acnid>\\d+)\\.(?P<metric2>.+?)\\.(?P<type>\\w+$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>(SmsDelivered|SmsUndelivered|TransmitToSmsPlatform|EnqueuedForSend|SmsSent|SmsCanceled|SmsDelayed|SmsFailed))\\.(?P<acncode>\\w+?)\\.(?P<cnrcode>\\w+?)\\.(?P<type>\\w+$) => ${metric1}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+?)\\.(?P<metric2>(size|queueSize|.+TransmittersCount)$) => ${metric1}.${metric2}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+?)\\.(?P<type>\\w+$) => ${metric1}",
			"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>[\\w-]+$) => ${metric1}",
		},
		"advisa": {
			"(?P<hostname>^avi\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>uptime$) => ${metric1}",
			"(?P<hostname>^avi\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+(Size|Speed|uptime)$) => ${metric1}",
			"(?P<hostname>^avi\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+)\\.(?P<type>.+$) => ${metric1}",
		},
	}
)

func BenchmarkParseReParser(b *testing.B) {
	for _, templates := range re_templates {
		p, _ := NewGraphiteReParser(".", "metric", templates, nil)

		for i := 0; i < b.N; i++ {
			p.ApplyTemplate("MfmsMonitor.manager-base-sbmanager3.zsbmng03.UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.priority.6.size")
		}
	}
}

func TestTemplateApplyReParser(t *testing.T) {
	var tests = map[string][]struct {
		input       string
		template    string
		measurement string
		tags        map[string]string
		err         string
	}{
		"sms": {
			{
				input:       "MfmsMonitor.avirouter-base-avirouter0.zchl09.AdvisaOutMessageCache.advisaOutMessageCachedMap.size",
				measurement: "AdvisaOutMessageCache.advisaOutMessageCachedMap.size",
				tags:        map[string]string{"component": "avirouter-base-avirouter0", "zone": "zchl09", "name": "avirouter0"},
			},
			{
				input:       "MfmsMonitor.imsichannel-megalabs-megalabs0.zchl04.ProtocolAdapter-agroros.undefinedMsisdnResponseMonitorAvgThroughputCounter",
				measurement: "ProtocolAdapter.undefinedMsisdnResponseMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "imsichannel-megalabs-megalabs0", "zone": "zchl04", "type": "agroros", "name": "megalabs0"},
			},
			{
				input:       "MfmsMonitor.manager-base-sbmanager3.zsbmng03.ComiConnectorOutMessageReceiver-sb8.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				measurement: "ComiConnectorOutMessageReceiver.peer.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "manager-base-sbmanager3", "zone": "zsbmng03", "peer": "sb8", "name": "sbmanager3"},
			},
			{
				input:       "MfmsMonitor.manager-base-manager0.zmng00.ComiConnectorOutMessageReceiver.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				measurement: "ComiConnectorOutMessageReceiver.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "manager-base-manager0", "zone": "zmng00", "name": "manager0"},
			},
			{
				input:       "MfmsMonitor.channel-smpp-alfacapmts0.zchl10.CmiChannelStateTransmitter-sbmanager1.channelStateProcessQueueProcessor.size",
				measurement: "CmiChannelStateTransmitter.peer.channelStateProcessQueueProcessor.size",
				tags:        map[string]string{"component": "channel-smpp-alfacapmts0", "zone": "zchl10", "peer": "sbmanager1", "name": "alfacapmts0"},
			},
			{
				input:       "MfmsMonitor.channel-beelineussd-beelineussd0.zchl04.CmiChannelStateTransmitter.channelStateProcessQueueProcessor.size",
				measurement: "CmiChannelStateTransmitter.channelStateProcessQueueProcessor.size",
				tags:        map[string]string{"component": "channel-beelineussd-beelineussd0", "zone": "zchl04", "name": "beelineussd0"},
			},
			{
				input:       "MfmsMonitor.connector-emailfileex-vtb2414.zcnr08.EmailFileExConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter.io",
				measurement: "ConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter.io",
				tags:        map[string]string{"component": "connector-emailfileex-vtb2414", "zone": "zcnr08", "name": "vtb2414"},
			},
			{
				input:       "MfmsMonitor.connector-alfa5-alfa15.zcnr00.AlfaConnector5DatabaseAccessor.connectorImsiResponseProcStatusAddQueueProcessor.size",
				measurement: "ConnectorDatabaseAccessor.connectorImsiResponseProcStatusAddQueueProcessor.size",
				tags:        map[string]string{"component": "connector-alfa5-alfa15", "zone": "zcnr00", "name": "alfa15"},
			},
			{
				input:       "MfmsMonitor.connector-fileex-russta3.zcnr00.FileExConnectorDatabaseAccessor.fileExSmsPostMessageAddQueueProcessor.size",
				measurement: "ConnectorDatabaseAccessor.SmsPostMessageAddQueueProcessor.size",
				tags:        map[string]string{"component": "connector-fileex-russta3", "zone": "zcnr00", "name": "russta3"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-autoins1.zcnr08.HpxConnector0DatabaseAccessor.hpxSmsGetMessageAddMonitorAvgSpeedCounter",
				measurement: "ConnectorDatabaseAccessor.SmsGetMessageAddMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "connector-hpx-autoins1", "zone": "zcnr08", "name": "autoins1"},
			},
			{
				input:       "MfmsMonitor.connectorbatch-smppimsi-bankofkazan0.zcnr03.SmppimsiConnectorBatchDatabaseAccessor.smppimsiSmsPostMessageAddMonitorAvgSpeedCounter",
				measurement: "ConnectorBatchDatabaseAccessor.SmsPostMessageAddMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "connectorbatch-smppimsi-bankofkazan0", "zone": "zcnr03", "name": "bankofkazan0"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb13.zcnr02.Gate.infobip0n1.CmiChannelInMessageTransmitterManager.channelInMessageProcessMonitorAvgSpeedCounter",
				measurement: "Gate.CmiChannelInMessageTransmitterManager.channelInMessageProcessMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "connector-sb1-sb13", "zone": "zcnr02", "gatecomponent": "infobip0n1", "name": "sb13"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb2.zsbcnr01.Gate.ifm.ws0n5.MonitorAccessor.monitorParameterProcessQueueProcessor.size",
				measurement: "Gate.MonitorAccessor.monitorParameterProcessQueueProcessor.size",
				tags:        map[string]string{"component": "connector-sb1-sb2", "zone": "zsbcnr01", "gatecomponent": "ifm.ws0n5", "name": "sb2"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb13.zcnr02.Gate.ermb0n0.ComiConnectorOutMessageTransmitter-manager1n0.connectorOutMessageProcessMonitorAvgSpeedCounter",
				measurement: "Gate.ComiConnectorOutMessageTransmitter.connectorOutMessageProcessMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "connector-sb1-sb13", "zone": "zcnr02", "gatecomponent": "ermb0n0", "peer": "manager1n0", "name": "sb13"},
			},
			{
				input:       "MfmsMonitor.receiver-base-receiver1.zchl04.RcoiConnectorInMessageTransmitter.binbank5.connectorInMessageProcessQueueProcessor.size",
				measurement: "RcoiConnectorInMessageTransmitter.peer.connectorInMessageProcessQueueProcessor.size",
				tags:        map[string]string{"component": "receiver-base-receiver1", "zone": "zchl04", "peer": "binbank5", "name": "receiver1"},
			},
			{
				input:       "MfmsMonitor.receiver-base-receiver0.zchl06.RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
				measurement: "RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
				tags:        map[string]string{"component": "receiver-base-receiver0", "zone": "zchl06", "name": "receiver0"},
			},
			{
				input:       "MfmsMonitor.receiver-base-receiver0.zchl06.RcoiConnectorInMessageTransmitter.ftc33.connectorInMessageProcessQueueProcessor.size",
				measurement: "RcoiConnectorInMessageTransmitter.peer.connectorInMessageProcessQueueProcessor.size",
				tags:        map[string]string{"component": "receiver-base-receiver0", "zone": "zchl06", "peer": "ftc33", "name": "receiver0"},
			},
			{
				input:       "MfmsMonitor.connector-emp-mospark1.zcnr03.ComiConnectorOutMessageTransmitterManager.processedConnectorOutMessageMonitorPercentileCounter.90",
				measurement: "ComiConnectorOutMessageTransmitterManager.processedConnectorOutMessageMonitorPercentileCounter",
				tags:        map[string]string{"component": "connector-emp-mospark1", "zone": "zcnr03", "percentile": "90", "name": "mospark1"},
			},
			{
				input:       "MfmsMonitor.manager-base-sbmanager3.zsbmng03.UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.priority.6.size",
				measurement: "UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.priority.size",
				tags:        map[string]string{"component": "manager-base-sbmanager3", "zone": "zsbmng03", "priority": "6", "name": "sbmanager3"},
			},
			{
				input:       "MfmsMonitor.manager-base-sbmanager3.zsbmng03.UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.size",
				measurement: "UndeliverableAddressChannelMessageProcessor.undeliverableAddressChannelMessageProcessQueueProcessor.size",
				tags:        map[string]string{"component": "manager-base-sbmanager3", "zone": "zsbmng03", "name": "sbmanager3"},
			},
			{
				input:       "MfmsMonitor.connector-chelin0-chelin0.zcnr04.Chelin0SmsPostMessageDlvStatusCache.chelin0SmsPostMessageDlvStatusBatchMap.size",
				measurement: "SmsPostMessageDlvStatusCache.SmsPostMessageDlvStatusBatchMap.size",
				tags:        map[string]string{"component": "connector-chelin0-chelin0", "zone": "zcnr04", "name": "chelin0"},
			},
			{
				input:       "MfmsMonitor.channel-smpp-alfacapmts0.zchl10.ResendProcessor.commandStatusMonitorAvgThroughputCounter.error",
				measurement: "ResendProcessor.commandStatusMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "channel-smpp-alfacapmts0", "zone": "zchl10", "status": "error", "name": "alfacapmts0"},
			},
			{
				input:       "MfmsMonitor.channel-smpp-beeline1.zchl06.DeliverSmsProcessor.channelInMessageProcessMonitorAvgThroughputCounter.79037676761",
				measurement: "DeliverSmsProcessor.channelInMessageProcessMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "channel-smpp-beeline1", "zone": "zchl06", "subject": "79037676761", "name": "beeline1"},
			},
			{
				input:       "MfmsMonitor.receiver-base-receiver0.zchl06.RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
				measurement: "RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
				tags:        map[string]string{"component": "receiver-base-receiver0", "zone": "zchl06", "name": "receiver0"},
			},
			{
				input:       "MfmsMonitor.manager-base-manager0.zmng00.ComiConnectorOutMessageReceiver-bistrodengi-web3.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				measurement: "ComiConnectorOutMessageReceiver.peer.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "manager-base-manager0", "zone": "zmng00", "peer": "bistrodengi-web3", "name": "manager0"},
			},
			{
				input:       "MfmsMonitor.connector-chelin0-chelin0.zcnr04.ChelinConnectorDatabaseAccessor.chelin0SmsPostBatchAddMonitorAvgSpeedCounter",
				measurement: "ConnectorDatabaseAccessor.SmsPostBatchAddMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "connector-chelin0-chelin0", "zone": "zcnr04", "name": "chelin0"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb4.zcnr02.Gate.ws1n1.ConnectorOutMessageEventProcessor.wsOutMessageQueue.SMS_UB_DVB.size",
				measurement: "Gate.ConnectorOutMessageEventProcessor.wsOutMessageQueue.size",
				tags:        map[string]string{"component": "connector-sb1-sb4", "zone": "zcnr02", "gatecomponent": "ws1n1", "queue": "SMS_UB_DVB", "name": "sb4"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb13.zcnr02.Gate.ws0n0.InMessageProcessor.messageQueue.sbbol_krim.size",
				measurement: "Gate.InMessageProcessor.messageQueue.size",
				tags:        map[string]string{"component": "connector-sb1-sb13", "zone": "zcnr02", "gatecomponent": "ws0n0", "queue": "sbbol_krim", "name": "sb13"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb4.zcnr02.Gate.mts1n0.DeliverStatusProcessor.deliverTimeCounter.1m",
				measurement: "Gate.DeliverStatusProcessor.deliverTimeCounter",
				tags:        map[string]string{"component": "connector-sb1-sb4", "zone": "zcnr02", "gatecomponent": "mts1n0", "time": "1m", "name": "sb4"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb13.zcnr02.Gate.infobip0n1.SendStatusProcessor.sendTimeCounter.more",
				measurement: "Gate.SendStatusProcessor.sendTimeCounter",
				tags:        map[string]string{"component": "connector-sb1-sb13", "zone": "zcnr02", "gatecomponent": "infobip0n1", "time": "more", "name": "sb13"},
			},
			{
				input:       "MfmsMonitor.imreceiver-base-imreceiver0.imsrv00.ImrvcoiConnectorInInstantMessageTransmitter.webclient0.connectorInInstantMessageProcessMonitorAvgSpeedCounter",
				measurement: "ImrvcoiConnectorInInstantMessageTransmitter.peer.connectorInInstantMessageProcessMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "imreceiver-base-imreceiver0", "zone": "imsrv00", "peer": "webclient0", "name": "imreceiver0"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb2.zsbcnr01.Gate.mts0n0.OutPacketProcessor.payloadOutPacketQueue.size",
				measurement: "Gate.OutPacketProcessor.payloadOutPacketQueue.size",
				tags:        map[string]string{"component": "connector-sb1-sb2", "zone": "zsbcnr01", "gatecomponent": "mts0n0", "name": "sb2"},
			},
			{
				input:       "MfmsMonitor.smppproxy-base-vtb24mts0.zchl06.OutPacketProcessor.payloadOutPacketQueue.vtb24mts0.size",
				measurement: "OutPacketProcessor.payloadOutPacketQueue.peer.size",
				tags:        map[string]string{"component": "smppproxy-base-vtb24mts0", "zone": "zchl06", "peer": "vtb24mts0", "name": "vtb24mts0"},
			},
			{
				input:       "MfmsMonitor.operatorprovider-base-operatorprovider0.zcnr02.login.lastRequestDate.magnit0",
				measurement: "login.lastRequestDate",
				tags:        map[string]string{"component": "operatorprovider-base-operatorprovider0", "zone": "zcnr02", "peer": "magnit0", "name": "operatorprovider0"},
			},
			{
				input:       "MfmsMonitor.imchannel-telegram-telegram0.imsrv02.telegramBot-RolfService24_7_bot.telegramInMessageDataProcessQueueProcessor.size",
				measurement: "telegramBot.telegramInMessageDataProcessQueueProcessor.size",
				tags:        map[string]string{"component": "imchannel-telegram-telegram0", "zone": "imsrv02", "bot": "RolfService24_7_bot", "name": "telegram0"},
			},
			{
				input:       "MfmsMonitor.smppproxy-base-mtsito0.zchl06.ClientProcessor.mtsito3.submitSmSequenceNumberMap.size",
				measurement: "ClientProcessor.submitSmSequenceNumberMap.size",
				tags:        map[string]string{"component": "smppproxy-base-mtsito0", "zone": "zchl06", "peer": "mtsito3", "name": "mtsito0"},
			},
			{
				input:       "MfmsMonitor.channel-smpp-sbmts8.zchl07.ChannelMonitorAccessor.deliveryMonitorDurationCounter.40000",
				measurement: "ChannelMonitorAccessor.deliveryMonitorDurationCounter",
				tags:        map[string]string{"component": "channel-smpp-sbmts8", "zone": "zchl07", "type": "40000", "name": "sbmts8"},
			},
			{
				input:       "MfmsMonitor.connector-sb1-sb4.zcnr02.Gate.route.AdminMBK_default.megafon0",
				measurement: "Gate.route",
				tags:        map[string]string{"component": "connector-sb1-sb4", "zone": "zcnr02", "system": "AdminMBK_default", "chl_group": "megafon0", "name": "sb4"},
			},
			{
				input:       "MfmsMonitor.manager-base-manager0.zmng00.ComiConnectorOutMessageReceiver-bistrodengi-web6.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				measurement: "ComiConnectorOutMessageReceiver.peer.receivedConnectorOutMessageMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "manager-base-manager0", "zone": "zmng00", "peer": "bistrodengi-web6", "name": "manager0"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-rbr1.zcnr06.n1.delivery.operator.kazakhstan-tele2.period.3600",
				measurement: "n1.delivery.operator.period",
				tags:        map[string]string{"component": "connector-hpx-rbr1", "zone": "zcnr06", "peer": "kazakhstan-tele2", "period": "3600", "name": "rbr1"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-rbr1.zcnr06.n1.delivery.operator.russia-ktktelecom.status.PROVIDER_UNDELIVERED",
				measurement: "n1.delivery.operator.status",
				tags:        map[string]string{"component": "connector-hpx-rbr1", "zone": "zcnr06", "peer": "russia-ktktelecom", "status": "PROVIDER_UNDELIVERED", "name": "rbr1"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-rbr1.zcnr06.n1.delivery.source.e-staff.period.43200",
				measurement: "n1.delivery.source.period",
				tags:        map[string]string{"component": "connector-hpx-rbr1", "zone": "zcnr06", "peer": "e-staff", "period": "43200", "name": "rbr1"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-rbr1.zcnr06.n1.delivery.source.stepUP.status.PROVIDER_FAILED",
				measurement: "n1.delivery.source.status",
				tags:        map[string]string{"component": "connector-hpx-rbr1", "zone": "zcnr06", "peer": "stepUP", "status": "PROVIDER_FAILED", "name": "rbr1"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-rbr1.zcnr06.n1.method.addPushDeviceMessages...requestsPerMinute",
				measurement: "n1.method.addPushDeviceMessages.requestsPerMinute",
				tags:        map[string]string{"component": "connector-hpx-rbr1", "zone": "zcnr06", "name": "rbr1"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-rbr1.zcnr06.n1.queue..SmppServerDlvEventQueueJob.itemsPerMinute",
				measurement: "n1.queue.SmppServerDlvEventQueueJob.itemsPerMinute",
				tags:        map[string]string{"component": "connector-hpx-rbr1", "zone": "zcnr06", "name": "rbr1"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-rbr1.zcnr06.n1.source.messagesPerMinute.RETAIL_MA",
				measurement: "n1.source.messagesPerMinute",
				tags:        map[string]string{"component": "connector-hpx-rbr1", "zone": "zcnr06", "peer": "RETAIL_MA", "name": "rbr1"},
			},
			{
				input:       "MfmsMonitor.smppproxy-base-mts1.zchl06.ClientProcessor.mts11.bound",
				measurement: "ClientProcessor.bound",
				tags:        map[string]string{"component": "smppproxy-base-mts1", "zone": "zchl06", "peer": "mts11", "name": "mts1"},
			},
			{
				input:       "MfmsMonitor.smppproxy-base-mts1.zchl06.OutPacketProcessor.payloadOutPacketQueue.mts19.size",
				measurement: "OutPacketProcessor.payloadOutPacketQueue.peer.size",
				tags:        map[string]string{"component": "smppproxy-base-mts1", "zone": "zchl06", "peer": "mts19", "name": "mts1"},
			},
			{
				input:       "MfmsMonitor.webclient-base-webclient0.zonline.CoimriConnectorOutInstantMessageTransmitter-imrouter0.connectorOutInstantMessageProcessQueueProcessor.size",
				measurement: "CoimriConnectorOutInstantMessageTransmitter.peer.connectorOutInstantMessageProcessQueueProcessor.size",
				tags:        map[string]string{"component": "webclient-base-webclient0", "zone": "zonline", "peer": "imrouter0", "name": "webclient0"},
			},
			{
				input:       "MfmsMonitor.avirouter-base-avirouter0.zchl09.ArcoiAdvisaOutMessageDlvEventTransmitter.advisaOutMessageDlvEventProcessMonitorAvgSpeedCounter",
				measurement: "ArcoiAdvisaOutMessageDlvEventTransmitter.advisaOutMessageDlvEventProcessMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "avirouter-base-avirouter0", "zone": "zchl09", "name": "avirouter0"},
			},
			{
				input:       "MfmsMonitor.avirouter-base-avirouter0.zchl09.ArcoiAdvisaOutMessageDlvEventTransmitter-alfa10.advisaOutMessageDlvEventProcessMonitorAvgSpeedCounter",
				measurement: "ArcoiAdvisaOutMessageDlvEventTransmitter.peer.advisaOutMessageDlvEventProcessMonitorAvgSpeedCounter",
				tags:        map[string]string{"component": "avirouter-base-avirouter0", "zone": "zchl09", "peer": "alfa10", "name": "avirouter0"},
			},
			{
				input:       "MfmsMonitor.connector-smpp-a31.zcnr05.SmppConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter",
				measurement: "ConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "connector-smpp-a31", "zone": "zcnr05", "name": "a31"},
			},
			{
				input:       "MfmsMonitor.connector-direct-testnccussd00.zcnr00.DirectConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter.io",
				measurement: "ConnectorDatabaseAccessor.databaseInteractionErrorMonitorAvgThroughputCounter.io",
				tags:        map[string]string{"component": "connector-direct-testnccussd00", "zone": "zcnr00", "name": "testnccussd00"},
			},
			{
				input:       "MfmsMonitor.channel-smpp-amdtelecom2.zchl07.ResendProcessor.commandStatusMonitorAvgThroughputCounter.error",
				measurement: "ResendProcessor.commandStatusMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "channel-smpp-amdtelecom2", "zone": "zchl07", "name": "amdtelecom2", "status": "error"},
			},
			{
				input:       "MfmsMonitor.connector-hpx-bancorp0.zcnr04.McoiConnectorOutMessageDlvEventReceiver.receivedConnectorOutMessageDlvEventMonitorAvgThroughputCounter.delivered",
				measurement: "McoiConnectorOutMessageDlvEventReceiver.receivedConnectorOutMessageDlvEventMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "connector-hpx-bancorp0", "zone": "zcnr04", "name": "bancorp0", "status": "delivered"},
			},
			{
				input:       "MfmsMonitor.channel-smpp-yandexmts0.zchl01.SubmitProcessor.deliveryDurationMonitorCounter.60",
				measurement: "SubmitProcessor.deliveryDurationMonitorCounter",
				tags:        map[string]string{"component": "channel-smpp-yandexmts0", "zone": "zchl01", "name": "yandexmts0", "type": "60"},
			},
		},
		"push": {
			{
				input:       "pushsrv00.connector-http-brooma_12.uptime",
				measurement: "uptime",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "connector-http-brooma_12"},
			},
			{
				input:       "pushsrv06.channel-push_gcmxmpp-prod_06.non-heap.committed",
				measurement: "non-heap",
				tags:        map[string]string{"hostname": "pushsrv06", "component": "channel-push_gcmxmpp-prod_06", "type": "committed"},
			},
			{
				input:       "pushsrv01.channel-push_wns-prod_01.CassandraSecurityTokenDataAccessor.errorCount.m1",
				measurement: "CassandraSecurityTokenDataAccessor.errorCount",
				tags:        map[string]string{"hostname": "pushsrv01", "component": "channel-push_wns-prod_01", "type": "m1"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.Transmitter.CmiPushNotification.channel-push_wp-prod_03.enqueueTimer.mean",
				measurement: "Transmitter.CmiPushNotification.enqueueTimer",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "type": "mean", "peer": "channel-push_wp-prod_03"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.Transmitter.CmiPushNotification.channel-push_gcmxmpp-prod_00.queueSize",
				measurement: "Transmitter.CmiPushNotification.queueSize",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "peer": "channel-push_gcmxmpp-prod_00"},
			},
			{
				input:       "pushdemo00.connector-advisa-demo_00.ClientOutMessageDlvTimeCounter.enqueuedToDelivered.p99",
				measurement: "ClientOutMessageDlvTimeCounter.enqueuedToDelivered",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-advisa-demo_00", "type": "p99"},
			},
			{
				input:       "pushdemo00.connector-advisa-demo_00.ClientOutMessageDlvTimeCounter.enqueuedToSent.mean",
				measurement: "ClientOutMessageDlvTimeCounter.enqueuedToSent",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-advisa-demo_00", "type": "mean"},
			},
			{
				input:       "pushsrv07.channel-push_apnshttp-prod_07.ApnsHttpChannelSender.ru_rsb_mobbank_334.successReverseGateMeter.count",
				measurement: "ApnsHttpChannelSender.successReverseGateMeter",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "channel-push_apnshttp-prod_07", "type": "count", "peer": "ru_rsb_mobbank_334"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.pools.Compressed-Class-Space.committed",
				measurement: "pools.Compressed-Class-Space",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "type": "committed"},
			},
			{
				input:       "pushsrv05.connector-http-tcsbank_03.ClientOutMessageSendService.commonSpeedMeter.m1",
				measurement: "ClientOutMessageSendService.commonSpeedMeter",
				tags:        map[string]string{"hostname": "pushsrv05", "component": "connector-http-tcsbank_03", "type": "m1"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.ClientOutMessageDlvStatusCounter.delivered.m1",
				measurement: "ClientOutMessageDlvStatusCounter",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "type": "m1", "status": "delivered"},
			},
			{
				input:       "pushsrv07.channel-push_gcmxmpp-prod_07.pools.Metaspace.max",
				measurement: "pools.Metaspace",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "channel-push_gcmxmpp-prod_07", "type": "max"},
			},
			{
				input:       "pushdemo00.channel-push_apnshttp-demo_00.CassandraDlvStatusInfoDataAccessor.getOutMessageDlvStatusByConnectorOutMessageIdTimer.p95",
				measurement: "CassandraDlvStatusInfoDataAccessor.getOutMessageDlvStatusByConnectorOutMessageIdTimer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "type": "p95"},
			},
			{
				input:       "pushdemo00.channel-push_gcm-demo_00.StoredQueue.1035.size",
				measurement: "StoredQueue.size",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_gcm-demo_00", "queue": "1035"},
			},
			{
				input:       "pushdemo00.connector-gate-sb_demo_01.requests.p99",
				measurement: "requests",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sb_demo_01", "type": "p99"},
			},
			{
				input:       "pushdemo00.connector-gate-sb_demo_01.retries-on-connection-error",
				measurement: "retries-on-connection-error",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sb_demo_01"},
			},
			{
				input:       "pushdemo00.channel-push_gcm-demo_00.GcmChannelSender.com_idamob_tinkoff_android_pro_1020.nonZeroNewMessageAvailableCount.count",
				measurement: "GcmChannelSender.nonZeroNewMessageAvailableCount",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_gcm-demo_00", "type": "count", "peer": "com_idamob_tinkoff_android_pro_1020"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.deviceStatusOutPacketFailSafeTransmitter.meter.count",
				measurement: "deviceStatusOutPacketFailSafeTransmitter.meter",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "type": "count"},
			},
			{
				input:       "pushdemo00.connector-http-brooma_demo_06.ConnectorOutMessageDlvStatusCounter.rejected.m1",
				measurement: "ConnectorOutMessageDlvStatusCounter",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-http-brooma_demo_06", "type": "m1", "status": "rejected"},
			},
			{
				input:       "pushdemo00.channel-push_apnshttp-demo_00.StoredQueue.Transmitter.CmiPushMessageDlvEvent.connector-http-tcsbankdemo_00.size",
				measurement: "StoredQueue.size",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "queue": "Transmitter.CmiPushMessageDlvEvent.connector-http-tcsbankdemo_00"},
			},
			{
				input:       "pushdemo00.connector-gate-sovcombankdemo_00.PlatformPacketListener.errorPacketCount",
				measurement: "PlatformPacketListener",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sovcombankdemo_00", "type": "errorPacketCount"}, //!!!
			},
			{
				input:       "pushdemo00.connector-gate-bcsdemo_00.PlatformPacketListener.outMessageTimer.p95",
				measurement: "PlatformPacketListener.outMessageTimer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-bcsdemo_00", "type": "p95"},
			},
			{
				input:       "pushdemo00.connector-gate-sbdemo_04.PlatformPacketListener.latestReceivedPacketTimestamp",
				measurement: "PlatformPacketListener",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sbdemo_04", "type": "latestReceivedPacketTimestamp"}, //!!!
			},
			{
				input:       "pushdemo00.channel-push_apnshttp-demo_00.ApnsHttpChannelSender.ru_sberbank_onlineiphone_beta_42.certificateValidDayCount",
				measurement: "ApnsHttpChannelSender.certificateValidDayCount",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "peer": "ru_sberbank_onlineiphone_beta_42"},
			},
			{
				input:       "pushsrv05.channel-push_apnshttp-prod_05.ApnsHttpChannelSender.com_brooma_threads-beta_377.nonZeroNewMessageAvailableCount.m1",
				measurement: "ApnsHttpChannelSender.nonZeroNewMessageAvailableCount",
				tags:        map[string]string{"hostname": "pushsrv05", "component": "channel-push_apnshttp-prod_05", "peer": "com_brooma_threads-beta_377", "type": "m1"},
			},
			{
				input:       "pushdemo00.connector-gate-openbankdemo_00.ClientMessageDatabaseProcessor.queueSize",
				measurement: "ClientMessageDatabaseProcessor.queueSize",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-openbankdemo_00"},
			},
			{
				input:       "pushsrv03.channel-sms_hpx-prod_01.HpxChannelProcessor-psbank.processTimer.m1",
				measurement: "HpxChannelProcessor.processTimer",
				tags:        map[string]string{"hostname": "pushsrv03", "component": "channel-sms_hpx-prod_01", "peer": "psbank", "type": "m1"},
			},
			{
				input:       "pushsrv00.server-web_push-prod_00.Transmitter.EnricherRequestCompletableFuture.sb.processTimer.p99",
				measurement: "Transmitter.EnricherRequestCompletableFuture.processTimer",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "server-web_push-prod_00", "peer": "sb", "type": "p99"},
			},
			{
				input:       "pushdemo00.connector-gate-vtb24demo_02.Transmitter.EnricherRequestCompletableFuture.vtb24.enqueueTimer.p999",
				measurement: "Transmitter.EnricherRequestCompletableFuture.enqueueTimer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-vtb24demo_02", "peer": "vtb24", "type": "p999"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.SmsCanceled.Sb.Base.count",
				measurement: "SmsCanceled",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "acncode": "Sb", "cnrcode": "Base", "type": "count"},
			},
			{
				input:       "pushsrv03.channel-sms_hpx-prod_01.HpxChannelProcessor-psbank.queueSize",
				measurement: "HpxChannelProcessor.queueSize",
				tags:        map[string]string{"hostname": "pushsrv03", "component": "channel-sms_hpx-prod_01", "peer": "psbank"},
			},
			{
				input:       "pushsrv03.connector-http-brooma_13.In_message_confirm.cnrId-131.acnId-17.count",
				measurement: "In_message_confirm",
				tags:        map[string]string{"hostname": "pushsrv03", "component": "connector-http-brooma_13", "cnrid": "131", "acnid": "17", "type": "count"},
			},
			{
				input:       "pushdemo00.server-web_push-demo_00.EnricherRequestRemoteServiceClient.skippedRequestMeter.count",
				measurement: "EnricherRequestRemoteServiceClient.skippedRequestMeter",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00", "type": "count"},
			},
			{
				input:       "pushdemo00.server-web_push-demo_00.EnricherRequestRemoteServiceClient.acnId-10.requestTimer.m1",
				measurement: "EnricherRequestRemoteServiceClient.requestTimer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00", "acnid": "10", "type": "m1"},
			},
			{
				input:       "pushsrv01.channel-sms_hpx-prod_00.SmsUndelivered.Sovcombank.Rosban.m1",
				measurement: "SmsUndelivered",
				tags:        map[string]string{"hostname": "pushsrv01", "component": "channel-sms_hpx-prod_00", "acncode": "Sovcombank", "cnrcode": "Rosban", "type": "m1"},
			},
			{
				input:       "pushsrv06.channel-sms_hpx-prod_02.StoredQueue.HpxChannel-vtb24-queue.size",
				measurement: "StoredQueue.size",
				tags:        map[string]string{"hostname": "pushsrv06", "component": "channel-sms_hpx-prod_02", "queue": "HpxChannel-vtb24-queue"},
			},
			{
				input:       "pushsrv05.channel-push_gcmxmpp-prod_05.StoredQueue.Transmitter.CmiPushMessageDlvEvent.connector-http-brooma_01.size",
				measurement: "StoredQueue.size",
				tags:        map[string]string{"hostname": "pushsrv05", "component": "channel-push_gcmxmpp-prod_05", "queue": "Transmitter.CmiPushMessageDlvEvent.connector-http-brooma_01"},
			},
			{
				input:       "pushsrv00.connector-http-brooma_12.SmsMessageTransmitterPool.errorCount.count",
				measurement: "SmsMessageTransmitterPool.errorCount",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "connector-http-brooma_12", "type": "count"},
			},
			{
				input:       "pushsrv00.connector-http-brooma_12.PushMessageTransmitterPool.aliveTransmittersCount",
				measurement: "PushMessageTransmitterPool.aliveTransmittersCount",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "connector-http-brooma_12"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.Transmitter.CmiSmsMessageDlvEvent.connector-http-tcsbank_02.enqueueTimer.mean",
				measurement: "Transmitter.CmiSmsMessageDlvEvent.enqueueTimer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "peer": "connector-http-tcsbank_02", "type": "mean"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.Transmitter.CmiSmsMessageDlvEvent.connector-db-test_demo.enqueueTimer.mean",
				measurement: "Transmitter.CmiSmsMessageDlvEvent.enqueueTimer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "peer": "connector-db-test_demo", "type": "mean"},
			},
			{
				input:       "pushdemo00.server-web_push-demo_00.Transmitter.CmiPushResendEvent.connector-http-broomademo_08.consumerAvailable",
				measurement: "Transmitter.CmiPushResendEvent.consumerAvailable",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00", "peer": "connector-http-broomademo_08"},
			},
		},
		"advisa": {
			{
				input:       "avitst00.advisa-database-monitor-test_00.DatabasePingService.avgBusyConnections.count",
				measurement: "DatabasePingService.avgBusyConnections",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "count"},
			},
			{
				input:       "avitst00.advisa-database-monitor-test_00.DataSourceFactoryBean.numIdle",
				measurement: "DataSourceFactoryBean",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "numIdle"},
			},
			{
				input:       "avitst00.advisa-database-monitor-test_00.non-heap.max",
				measurement: "non-heap",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "max"},
			},
			{
				input:       "avitst00.advisa-database-monitor-test_00.pools.PS-Survivor-Space.used-after-gc",
				measurement: "pools.PS-Survivor-Space",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "used-after-gc"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.advisaOperationProcessor.retryQueueSize",
				measurement: "advisaOperationProcessor.retryQueueSize",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.coasConnectorOperationDlvEventConsumer.enqueueTimer.p999",
				measurement: "coasConnectorOperationDlvEventConsumer.enqueueTimer",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00", "type": "p999"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.CoasSubscriptionStatusEventTransmitterPool.aliveTransmittersCount",
				measurement: "CoasSubscriptionStatusEventTransmitterPool",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00", "type": "aliveTransmittersCount"},
			},
			{
				input:       "avitst00.connector-terminal_location_aggregator-test_00.StoredQueue.terminalLocationsAggregateMultiPointProcessorRetry.size",
				measurement: "StoredQueue.terminalLocationsAggregateMultiPointProcessorRetry",
				tags:        map[string]string{"component": "connector-terminal_location_aggregator-test_00", "hostname": "avitst00", "type": "size"},
			},
			{
				input:       "avitst00.server-web_advisa-test_00.uptime",
				measurement: "uptime",
				tags:        map[string]string{"component": "server-web_advisa-test_00", "hostname": "avitst00"},
			},
		},
	}

	for system, templates := range re_templates {
		fmt.Println("!!! " + system)
		p, err := NewGraphiteReParser(".", "metric", templates, nil)
		if err != nil {
			t.Fatal("error parsin regexp: ", err)
		}
		for _, test := range tests[system] {
			measurement, tags, _, _ := p.ApplyTemplate(test.input)
			fmt.Println(measurement, tags)
			if measurement != test.measurement {
				t.Fatalf("name parse failed. expected %v got %v", test.measurement, measurement)
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
}
