package graphite

import (
	"fmt"
	"testing"
)

var (
	re_templates []string = []string{
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>Gate\\.route)\\.(?P<system>[\\w-]+?)\\.(?P<chl_group>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<measurement>.+?TimeCounter)\\.(?P<time>\\w+?$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<measurement>.+?[Mm]essageQueue)\\.(?P<queue>\\w+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<measurement>\\w+?Transmitter)-(?P<peer>\\w+?\\d+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>Gate)\\.(?P<gatecomponent>(ifm\\.\\w+?)|(\\w+?))\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<measurement>Connector(Batch)?)(?:\\d*?)(?P<_measurement>DatabaseAccessor)\\.(?:\\w*?)(?P<measurement>(SmsPost|SmsGet).+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<measurement>Connector)(?:\\d*?)(?P<_measurement>DatabaseAccessor)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?:\\w*?)(?P<measurement>(SmsPost|SmsGet)\\w+)\\.(?:\\w*?)(?P<measurement>(SmsPost|SmsGet).+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?[Mm]essageQueue)\\.(?P<queue>\\w+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>RcoiConnectorInMessageTransmitter)\\.(?P<peer>\\w+?\\d+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>telegramBot)-(?P<bot>[\\w-]+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>ClientProcessor)\\.(?P<peer>[\\w]+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>\\w+?Transmitter)(?:-|\\.)(?P<peer>\\w+?\\d+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>\\w+?Receiver)-(?P<peer>[\\w-]+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>\\w+?Adapter)-(?P<type>[\\w]+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?deliveryMonitorDurationCounter)\\.(?P<type>\\w+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?commandStatusMonitorAvgThroughputCounter)\\.(?P<status>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?payloadOutPacketQueue)\\.(?P<peer>[\\w]+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?channelInMessageProcessMonitorAvgThroughputCounter)\\.(?P<subject>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?PercentileCounter)\\.(?P<percentile>\\w+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?(lastRequestDate|lastMessageTime|messagesPerMinute))\\.(?P<peer>\\w+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?(operator|source))\\.(?P<peer>[\\w-]+)\\.(?P<measurement>period)\\.(?P<period>\\w+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?(operator|source))\\.(?P<peer>[\\w-]+)\\.(?P<measurement>status)\\.(?P<status>\\w+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?\\.priority)\\.(?P<priority>\\w+?)\\.(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+?)\\.(?:\\.+)(?P<measurement>.+$)",
		"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<measurement>.+$)",

		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>\\w+?ChannelSender)(?:-appId-)(?P<appid>\\d+?)\\.(?P<measurement>.+?(?:Count|Timer))\\.(?P<type>\\w+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>\\w+?ChannelSender)(?:-appId-)(?P<appid>\\d+?)\\.(?P<measurement>.+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>\\w+?ChannelSender)\\.(?P<appid>[\\w.-]+?)\\.(?P<measurement>(?:http|send|success)[\\w]+?)\\.(?P<type>\\w+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>\\w+?OutMessageDlvStatusCounter)\\.(?P<status>\\w+?)\\.(?P<type>\\w+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>.+?(Delivered|Sent|Failed|Count|Timer|Meter|\\.meter))\\.(?P<type>\\w+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>pools)\\.(?P<measurement>[\\w-]+?)\\.(?P<type>\\w+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>StoredQueue)\\.(?P<queue>[\\w]+?)\\.(?P<measurement>\\w+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>requests)\\.(?P<type>\\w+$)",
		"(?P<hostname>^push\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<measurement>.+$)",

		"(?P<component>[\\w-]+)-(?P<node>[\\w-]+?_\\d{2})\\.(?P<measurement>.+?(Transmitter|Invoker))\\.(?P<peer>[\\w-]+?)\\.(?P<measurement>timer)\\.(?P<type>.+$)",
		"(?P<component>[\\w-]+)-(?P<node>[\\w-]+?_\\d{2})\\.(?P<measurement>.+?Transmitter)\\.(?P<peer>[\\w-]+?)\\.(?P<measurement>.+$)",
		"(?P<component>[\\w-]+)-(?P<node>[\\w-]+?_\\d{2})\\.(?P<measurement>.+?Pool)\\.(?P<peer>[\\w-]+?)\\.(?P<measurement>.+$)",
		"(?P<component>[\\w-]+)-(?P<node>[\\w-]+?_\\d{2})\\.(?P<measurement>.+?(timer|histogram|meter|Time|Timer|Meter|Count|Duration|Locations))\\.(?P<type>.+$)",
		"(?P<component>[\\w-]+)-(?P<node>[\\w-]+?_\\d{2})\\.(?P<measurement>.+?)\\.(?P<proto>/.+)\\.(?P<measurement>SystemError)\\.(?P<type>.+$)",
		"(?P<component>[\\w-]+)-(?P<node>[\\w-]+?_\\d{2})\\.(?P<measurement>.+$)",
	}
)

func BenchmarkParseReParser(b *testing.B) {
	p, _ := NewGraphiteReParser(".", "measurement", re_templates, nil)

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
			measurement: "ComiConnectorOutMessageReceiver.receivedConnectorOutMessageMonitorAvgThroughputCounter",
			tags:        map[string]string{"component": "manager-base-sbmanager3", "zone": "zsbmng03", "peer": "sb8", "name": "sbmanager3"},
		},
		{
			input:       "MfmsMonitor.channel-smpp-alfacapmts0.zchl10.CmiChannelStateTransmitter-sbmanager1.channelStateProcessQueueProcessor.size",
			measurement: "CmiChannelStateTransmitter.channelStateProcessQueueProcessor.size",
			tags:        map[string]string{"component": "channel-smpp-alfacapmts0", "zone": "zchl10", "peer": "sbmanager1", "name": "alfacapmts0"},
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
			measurement: "RcoiConnectorInMessageTransmitter.connectorInMessageProcessQueueProcessor.size",
			tags:        map[string]string{"component": "receiver-base-receiver1", "zone": "zchl04", "peer": "binbank5", "name": "receiver1"},
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
			measurement: "ComiConnectorOutMessageReceiver.receivedConnectorOutMessageMonitorAvgThroughputCounter",
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
			measurement: "ImrvcoiConnectorInInstantMessageTransmitter.connectorInInstantMessageProcessMonitorAvgSpeedCounter",
			tags:        map[string]string{"component": "imreceiver-base-imreceiver0", "zone": "imsrv00", "peer": "webclient0", "name": "imreceiver0"},
		},
		{
			input:       "MfmsMonitor.connector-sb1-sb2.zsbcnr01.Gate.mts0n0.OutPacketProcessor.payloadOutPacketQueue.size",
			measurement: "Gate.OutPacketProcessor.payloadOutPacketQueue.size",
			tags:        map[string]string{"component": "connector-sb1-sb2", "zone": "zsbcnr01", "gatecomponent": "mts0n0", "name": "sb2"},
		},
		{
			input:       "MfmsMonitor.smppproxy-base-vtb24mts0.zchl06.OutPacketProcessor.payloadOutPacketQueue.vtb24mts0.size",
			measurement: "OutPacketProcessor.payloadOutPacketQueue.size",
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
			measurement: "ComiConnectorOutMessageReceiver.receivedConnectorOutMessageMonitorAvgThroughputCounter",
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
		// =====================================================================================================================
		// =====================================================================================================================
		// =====================================================================================================================

		{
			input:       "pushdemo00.server-web_push-demo_00.non-heap.committed",
			measurement: "non-heap.committed",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00"},
		},
		{
			input:       "pushdemo00.channel-push_gcm-demo_00.CassandraSecurityTokenDataAccessor.errorCount.m15_rate",
			measurement: "CassandraSecurityTokenDataAccessor.errorCount",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_gcm-demo_00", "type": "m15_rate"},
		},
		{
			input:       "pushdemo00.channel-push_apnshttp-demo_00.ApnsHttpChannelSender-appId-1003.dequeueTimer.max",
			measurement: "ApnsHttpChannelSender.dequeueTimer",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "type": "max", "appid": "1003"},
		},
		{
			input:       "pushdemo00.channel-push_apnshttp-demo_00.ApnsHttpChannelSender-appId-1003.queueSize",
			measurement: "ApnsHttpChannelSender.queueSize",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "appid": "1003"},
		},
		{
			input:       "pushdemo00.connector-gate-openbank_demo_01.ClientOutMessageDlvTimeCounter.enqueudToDelivered.p50",
			measurement: "ClientOutMessageDlvTimeCounter.enqueudToDelivered",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-openbank_demo_01", "type": "p50"},
		},
		{
			input:       "pushdemo00.connector-gate-openbank_demo_00.ClientOutMessageDlvTimeCounter.enqueudToSent.m5_rate",
			measurement: "ClientOutMessageDlvTimeCounter.enqueudToSent",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-openbank_demo_00", "type": "m5_rate"},
		},
		{
			input:       "pushdemo00.channel-push_apnshttp-demo_00.ApnsHttpChannelSender.com.advisa.advisaenterprise.vtb.201.successReverseGateMeter.m1_rate",
			measurement: "ApnsHttpChannelSender.successReverseGateMeter",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "type": "m1_rate", "appid": "com.advisa.advisaenterprise.vtb.201"},
		},
		{
			input:       "pushdemo00.channel-push_apnshttp-demo_00.ApnsHttpChannelSender.ru.mfms.push-test.voip.310.successReverseGateMeter.mean_rate",
			measurement: "ApnsHttpChannelSender.successReverseGateMeter",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "type": "mean_rate", "appid": "ru.mfms.push-test.voip.310"},
		},
		{
			input:       "pushdemo00.server-web_push-demo_00.pools.Compressed-Class-Space.committed",
			measurement: "pools.Compressed-Class-Space",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00", "type": "committed"},
		},
		{
			input:       "pushdemo00.connector-httpxml_securitytoken-test_demo_00.ClientOutMessageSendService.commonSpeedMeter.mean_rate",
			measurement: "ClientOutMessageSendService.commonSpeedMeter",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-httpxml_securitytoken-test_demo_00", "type": "mean_rate"},
		},
		{
			input:       "pushdemo00.connector-httpxml_securitytoken-test_demo_00.ClientOutMessageDlvStatusCounter.anyDelivered.m15_rate",
			measurement: "ClientOutMessageDlvStatusCounter",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-httpxml_securitytoken-test_demo_00", "type": "m15_rate", "status": "anyDelivered"},
		},
		{
			input:       "pushdemo00.channel-sms_hpx-demo_00.pools.Metaspace.max",
			measurement: "pools.Metaspace",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "type": "max"},
		},
		{
			input:       "pushdemo00.channel-push_apnshttp-demo_00.CassandraDlvStatusInfoDataAccessor.getOutMessageDlvStatusByConnectorOutMessageIdAndAccountIdTimer.m15_rate",
			measurement: "CassandraDlvStatusInfoDataAccessor.getOutMessageDlvStatusByConnectorOutMessageIdAndAccountIdTimer",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "type": "m15_rate"},
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
			input:       "pushdemo00.channel-push_gcm-demo_00.GcmChannelSender-appId-1014.processTimer.p95",
			measurement: "GcmChannelSender.processTimer",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_gcm-demo_00", "type": "p95", "appid": "1014"},
		},
		{
			input:       "pushdemo00.connector-gate-sb_demo_04.deviceStatusOutPacketFailSafeTransmitter.meter.samples",
			measurement: "deviceStatusOutPacketFailSafeTransmitter.meter",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sb_demo_04", "type": "samples"},
		},
		{
			input:       "pushdemo00.connector-http-brooma_demo_06.ConnectorOutMessageDlvStatusCounter.rejected.m1_rate",
			measurement: "ConnectorOutMessageDlvStatusCounter",
			tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-http-brooma_demo_06", "type": "m1_rate", "status": "rejected"},
		},
		// =====================================================================================================================
		// =====================================================================================================================
		// =====================================================================================================================

		{
			input:       "connector-currencycourse_yahoo-prod_00.DataSourceFactoryBean.numActive",
			measurement: "DataSourceFactoryBean.numActive",
			tags:        map[string]string{"component": "connector-currencycourse_yahoo", "node": "prod_00"},
		},
		{
			input:       "connector-storage_s3-prod_00.pools.Compressed-Class-Space.usage",
			measurement: "pools.Compressed-Class-Space.usage",
			tags:        map[string]string{"component": "connector-storage_s3", "node": "prod_00"},
		},
		{
			input:       "connector-operations_smsparse-prod_00.AdvisaOperationDatabaseProcessor.StoredQueueProcessorJob.timer.max",
			measurement: "AdvisaOperationDatabaseProcessor.StoredQueueProcessorJob.timer",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_00", "type": "max"},
		},
		{
			input:       "server-web_advisa-prod_01.CoasAdvisaSubscriptionStatusEventConsumerImpl.StoredQueueProcessorJob.meter.m5_rate",
			measurement: "CoasAdvisaSubscriptionStatusEventConsumerImpl.StoredQueueProcessorJob.meter",
			tags:        map[string]string{"component": "server-web_advisa", "node": "prod_01", "type": "m5_rate"},
		},
		{
			input:       "server-web_advisa-prod_01.CoasBankNotificationConsumerImpl.StoredQueueProcessorJob.histogram.p98",
			measurement: "CoasBankNotificationConsumerImpl.StoredQueueProcessorJob.histogram",
			tags:        map[string]string{"component": "server-web_advisa", "node": "prod_01", "type": "p98"},
		},
		{
			input:       "server-web_advisa-prod_01.CoasCardTransactionWithNewTerminalDatabaseProcessor.createTerminalTimer.min",
			measurement: "CoasCardTransactionWithNewTerminalDatabaseProcessor.createTerminalTimer",
			tags:        map[string]string{"component": "server-web_advisa", "node": "prod_01", "type": "min"},
		},
		{
			input:       "server-web_advisa-prod_01.SystemJournalService.writeFirstSyncIfNeededDuration.mean_rate",
			measurement: "SystemJournalService.writeFirstSyncIfNeededDuration",
			tags:        map[string]string{"component": "server-web_advisa", "node": "prod_01", "type": "mean_rate"},
		},
		{
			input:       "connector-operations_smsparse-prod_01.AdvisaOperationDlvEventDatabaseProcessor.StoredQueueProcessorJob.histogram.stddev",
			measurement: "AdvisaOperationDlvEventDatabaseProcessor.StoredQueueProcessorJob.histogram",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_01", "type": "stddev"},
		},
		{
			input:       "connector-operations_smsparse-prod_01.BankIncomingSmsService.avgBatchTime.mean_rate",
			measurement: "BankIncomingSmsService.avgBatchTime",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_01", "type": "mean_rate"},
		},
		{
			input:       "connector-operations_smsparse-prod_00.AdvisaConnectorSubscriptionDatabaseAccessor.databaseErrorCount.mean_rate",
			measurement: "AdvisaConnectorSubscriptionDatabaseAccessor.databaseErrorCount",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_00", "type": "mean_rate"},
		},
		{
			input:       "connector-operations_smsparse-prod_00.AdvisaOperationProcessor.avgSpeedMeter.mean_rate",
			measurement: "AdvisaOperationProcessor.avgSpeedMeter",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_00", "type": "mean_rate"},
		},
		{
			input:       "connector-operations_smsparse-prod_00.CoasBonusBallUpdateOperationTransmitterPool.server-web_advisa-prod_01.percent",
			measurement: "CoasBonusBallUpdateOperationTransmitterPool.percent",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_00", "peer": "server-web_advisa-prod_01"},
		},
		{
			input:       "connector-operations_smsparse-prod_00.CoasSubscriptionStatusEventTransmitterPool.Transmitter.server-web_advisa-prod_00.timer.p95",
			measurement: "CoasSubscriptionStatusEventTransmitterPool.Transmitter.timer",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_00", "peer": "server-web_advisa-prod_00", "type": "p95"},
		},
		{
			input:       "connector-operations_smsparse-prod_00.CoasAccountTransactionTransmitterPool.Transmitter.server-web_advisa-prod_01.queueSize",
			measurement: "CoasAccountTransactionTransmitterPool.Transmitter.queueSize",
			tags:        map[string]string{"component": "connector-operations_smsparse", "node": "prod_00", "peer": "server-web_advisa-prod_01"},
		},
		{
			input:       "connector-registration-prod_00.CoasAdvisaRegistrationConfirmationResponseTransmitter.server-web_advisa-prod_00.timer.p999",
			measurement: "CoasAdvisaRegistrationConfirmationResponseTransmitter.timer",
			tags:        map[string]string{"component": "connector-registration", "node": "prod_00", "peer": "server-web_advisa-prod_00", "type": "p999"},
		},
		{
			input:       "server-web_advisa-prod_00.com.mfms.advisa.web.utils.spring.ProtobufHttpMessageConverter./service/sync/do.ru.raiffeisenbank.rconnect.IOS.SystemError.m1_rate",
			measurement: "com.mfms.advisa.web.utils.spring.ProtobufHttpMessageConverter.SystemError",
			tags:        map[string]string{"component": "server-web_advisa", "node": "prod_00", "proto": "/service/sync/do.ru.raiffeisenbank.rconnect.IOS", "type": "m1_rate"},
		},
		{
			input:       "connector-terminal-prod_00.TerminalFullTextIndexingProcessor.terminalGetWithLocations.p95",
			measurement: "TerminalFullTextIndexingProcessor.terminalGetWithLocations",
			tags:        map[string]string{"component": "connector-terminal", "node": "prod_00", "type": "p95"},
		},
		{
			input:       "connector-terminal_enricher-prod_00.CoasTerminalConnectorSimilarServicePool.connector-terminal-prod_01.percent",
			measurement: "CoasTerminalConnectorSimilarServicePool.percent",
			tags:        map[string]string{"component": "connector-terminal_enricher", "node": "prod_00", "peer": "connector-terminal-prod_01"},
		},
		{
			input:       "connector-terminal_enricher-prod_00.CoasTerminalConnectorSimilarServicePool.ServiceInvoker.connector-terminal-prod_01.timer.p95",
			measurement: "CoasTerminalConnectorSimilarServicePool.ServiceInvoker.timer",
			tags:        map[string]string{"component": "connector-terminal_enricher", "node": "prod_00", "peer": "connector-terminal-prod_01", "type": "p95"},
		},
		{
			input:       "server-problem_devices-prod_00.G1-Old-Generation.count",
			measurement: "G1-Old-Generation.count",
			tags:        map[string]string{"component": "server-problem_devices", "node": "prod_00"},
		},
		{
			input:       "server-web_advisa-prod_01.heap.used",
			measurement: "heap.used",
			tags:        map[string]string{"component": "server-web_advisa", "node": "prod_01"},
		},
		{
			input:       "server-push_reports-prod_00.pools.G1-Survivor-Space.usage",
			measurement: "pools.G1-Survivor-Space.usage",
			tags:        map[string]string{"component": "server-push_reports", "node": "prod_00"},
		},
	}

	p, err := NewGraphiteReParser(".", "measurement", re_templates, nil)
	if err != nil {
		t.Fatal("error parsin regexp: ", err)
	}

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
