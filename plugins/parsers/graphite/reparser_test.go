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
			"(?:^MfmsMonitor)\\.(?P<component>\\w+?-\\w+?-(?P<name>\\w+?\\d+))\\.(?P<zone>\\w+?)\\.(?P<metric1>\\w+)\\.(?P<status>\\w+)(?P<metric2>ViberRespMonitorAvgThroughputCounter$) => ${metric1}.${metric2}",
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
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+(enqueuedToDelivered|enqueuedToSent|Time|Timer|Timer\\.\\w+))\\.(?P<type>[\\w-]+$) => timer",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+(statusMeter|StatusCounter))\\.(?P<status>[\\w-]+)\\.(?P<type>[\\w-]+$) => counter.status",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+(Counter\\.[\\w.-]+|Count|Meter|meter|enqueued|Available))\\.(?P<type>[\\w-]+$) => counter",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+Pool)\\.(?P<type>[\\w-]+Count$) => pool",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+Pool\\.[\\w.-]+Count)\\.(?P<type>[\\w-]+$) => pool.counter",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>StoredQueue\\.Transmitter\\.[\\w.-]+)\\.(?P<type>[\\w-]+$) => storedqueue.transmitter",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>StoredQueue\\.[\\w.-]+)\\.(?P<type>[\\w-]+$) => storedqueue",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>StoredMap\\.[\\w.-]+)\\.(?P<type>[\\w-]+$) => storedmap",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>queueSize$) => queue",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>(EnqueuedForSend|SmsCanceled|SmsDelivered|SmsUndelivered|TransmitToSmsPlatform)[\\w.-]+)\\.(?P<type>[\\w-]+$) => counter",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>In_message[\\w.-]+)\\.(?P<type>count|m1$) => counter",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>latest[\\w-]+Timestamp$) => timestamp",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>(cacheSize|hitCount|missCount|successRate)$) => cache",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>idCacheSize$) => idcache",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]*?(heap|non-heap|-Cache|-Space|-Gen|Metaspace|total))\\.(?P<type>[\\w-]+$) => memory",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]*?(Generation|MarkSweep|Scavenge))\\.(?P<type>[\\w-]+$) => gc",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>(certificateRevoked|pnsAvailable|sslWorking|bound|consumerAvailable)$) => state",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+Pool)\\.(?P<type>[\\w-]+Count$) => pool",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+Pool\\.[\\w.-]+Count)\\.(?P<type>[\\w-]+$) => pool.counter",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>(errorPacketCount|certificateValidDayCount|outQueueSize|xmppQueueSize)$) => value",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>(protocolVersion|notificationsPushed|notificationsPushedWithError)$) => value",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+)\\.(?P<type>(disconnectCount|notificationsDelivered|notificationsEnqueued|notificationsUnknown|descriptors)$) => value",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>[\\w.-]+Bean)\\.(?P<type>[\\w-]+$) => bean",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.requests\\.(?P<type>[\\w-]+$) => cassandra.requests",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<midname>bytes-received|bytes-sent)\\.(?P<type>[\\w-]+$) => cassandra.counter",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.(?P<type>((\\w+-[\\w-]+)|ignores|retries|unavailables)$) => cassandra",
			"(?P<hostname>^push\\w+)\\.(?P<component>\\w+-\\w+-(?P<name>\\w+))\\.uptime => uptime",
		},
		"advisa": {
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>StoredQueue\\.Transmitter\\.[\\w.-]+?)\\.(?P<type>[\\w-]+$) => storedqueue.transmitter",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>StoredQueue\\.[\\w.-]+?)\\.(?P<type>[\\w-]+$) => storedqueue",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>Transmitter\\.[\\w.-]+?(Time|Timer|Override|Operations|Merchant|Terminal|DataSource|QueryExecution|OpeningConnection|Duration|Locations))\\.(?P<type>[\\w-]+$) => transmitter.timer",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>Transmitter\\.[\\w.-]+?(Count|Meter|Imports|Connections))\\.(?P<type>[\\w-]+$) => transmitter.counter",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>Transmitter\\.[\\w.-]+?)\\.(?P<type>[\\w-]+$) => transmitter",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>StoredMap\\.[\\w.-]+?)\\.(?P<type>[\\w-]+$) => storedmap",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?(Count|Meter|Imports|Connections))\\.(?P<type>[\\w-]+$) => counter",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?Bean)\\.(?P<type>[\\w-]+$) => bean",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]*?(heap|non-heap|-Cache|-Space|-Gen|Metaspace|total))\\.(?P<type>[\\w-]+$) => memory",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]*?(Generation|MarkSweep|Scavenge))\\.(?P<type>[\\w-]+$) => gc",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?(Time|Timer|timer|Override|Operations|Merchant|Terminal|DataSource|QueryExecution|OpeningConnection|Duration|Locations))\\.(?P<type>[\\w-]+$) => timer",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?Pool)\\.(?P<type>[\\w-]+$) => pool",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?IdCache)\\.(?P<type>[\\w-]+$) => idcache",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?Cache)\\.(?P<type>[\\w-]+$) => cache",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?IndexService)\\.(?P<type>[\\w-]+$) => index",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?)\\.(?P<type>(queueSize|retryQueueSize)$) => queue",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?)\\.(?P<type>(notIndexedCount|descriptors)$) => value",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<midname>[\\w.-]+?)\\.(?P<type>[\\w-]+Speed$) => speed",
			"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.uptime$ => uptime",
			//"(?P<hostname>^\\w+?)\\.(?P<component>[\\w-]+?\\d+)\\.(?P<metric1>.+$) => ${metric1}",
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
			{
				input:       "MfmsMonitor.imchannel-viber-viber0.imsrv00.ImChannelOutInstantMessageProcessor.noSuitableDeviceViberRespMonitorAvgThroughputCounter",
				measurement: "ImChannelOutInstantMessageProcessor.ViberRespMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "imchannel-viber-viber0", "zone": "imsrv00", "name": "viber0", "status": "noSuitableDevice"},
			},
			{
				input:       "MfmsMonitor.imchannel-viber-viber1.imsrv01.ImChannelOutInstantMessageProcessor.badParametersViberRespMonitorAvgThroughputCounter",
				measurement: "ImChannelOutInstantMessageProcessor.ViberRespMonitorAvgThroughputCounter",
				tags:        map[string]string{"component": "imchannel-viber-viber1", "zone": "imsrv01", "name": "viber1", "status": "badParameters"},
			},
		},
		"push": {
			{
				input:       "pushsrv00.connector-http-brooma_12.uptime",
				measurement: "uptime",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "connector-http-brooma_12", "name": "brooma_12"},
			},
			{
				input:       "pushsrv06.channel-push_gcmxmpp-prod_06.non-heap.committed",
				measurement: "memory",
				tags:        map[string]string{"hostname": "pushsrv06", "component": "channel-push_gcmxmpp-prod_06", "name": "prod_06", "type": "committed", "midname": "non-heap"},
			},
			{
				input:       "pushsrv01.channel-push_wns-prod_01.CassandraSecurityTokenDataAccessor.errorCount.m1",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv01", "component": "channel-push_wns-prod_01", "name": "prod_01", "type": "m1", "midname": "CassandraSecurityTokenDataAccessor.errorCount"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.Transmitter.CmiPushNotification.channel-push_wp-prod_03.enqueueTimer.mean",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "name": "sb_06", "type": "mean", "midname": "Transmitter.CmiPushNotification.channel-push_wp-prod_03.enqueueTimer"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.Transmitter.CmiPushNotification.channel-push_gcmxmpp-prod_00.queueSize",
				measurement: "queue",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "name": "sb_06", "type": "queueSize", "midname": "Transmitter.CmiPushNotification.channel-push_gcmxmpp-prod_00"},
			},
			{
				input:       "pushdemo00.connector-advisa-demo_00.ClientOutMessageDlvTimeCounter.enqueuedToDelivered.p99",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-advisa-demo_00", "name": "demo_00", "type": "p99", "midname": "ClientOutMessageDlvTimeCounter.enqueuedToDelivered"},
			},
			{
				input:       "pushdemo00.connector-advisa-demo_00.ClientOutMessageDlvTimeCounter.enqueuedToSent.mean",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-advisa-demo_00", "name": "demo_00", "type": "mean", "midname": "ClientOutMessageDlvTimeCounter.enqueuedToSent"},
			},
			{
				input:       "pushsrv07.channel-push_apnshttp-prod_07.ApnsHttpChannelSender.ru_rsb_mobbank_334.successReverseGateMeter.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "channel-push_apnshttp-prod_07", "name": "prod_07", "type": "count", "midname": "ApnsHttpChannelSender.ru_rsb_mobbank_334.successReverseGateMeter"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.pools.Compressed-Class-Space.committed",
				measurement: "memory",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "name": "sb_06", "type": "committed", "midname": "pools.Compressed-Class-Space"},
			},
			{
				input:       "pushsrv05.connector-http-tcsbank_03.ClientOutMessageSendService.commonSpeedMeter.m1",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv05", "component": "connector-http-tcsbank_03", "name": "tcsbank_03", "type": "m1", "midname": "ClientOutMessageSendService.commonSpeedMeter"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.ClientOutMessageDlvStatusCounter.delivered.m1",
				measurement: "counter.status",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "name": "sb_06", "type": "m1", "midname": "ClientOutMessageDlvStatusCounter", "status": "delivered"},
			},
			{
				input:       "pushsrv07.channel-push_gcmxmpp-prod_07.pools.Metaspace.max",
				measurement: "memory",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "channel-push_gcmxmpp-prod_07", "name": "prod_07", "type": "max", "midname": "pools.Metaspace"},
			},
			{
				input:       "pushdemo00.channel-push_apnshttp-demo_00.CassandraDlvStatusInfoDataAccessor.getOutMessageDlvStatusByConnectorOutMessageIdTimer.p95",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "name": "demo_00", "type": "p95", "midname": "CassandraDlvStatusInfoDataAccessor.getOutMessageDlvStatusByConnectorOutMessageIdTimer"},
			},
			{
				input:       "pushdemo00.channel-push_gcm-demo_00.StoredQueue.1035.size",
				measurement: "storedqueue",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_gcm-demo_00", "name": "demo_00", "type": "size", "midname": "StoredQueue.1035"},
			},
			{
				input:       "pushdemo00.connector-gate-sb_demo_01.requests.p99",
				measurement: "cassandra.requests",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sb_demo_01", "name": "sb_demo_01", "type": "p99"},
			},
			{
				input:       "pushdemo00.connector-gate-sb_demo_01.retries-on-connection-error",
				measurement: "cassandra",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sb_demo_01", "name": "sb_demo_01", "type": "retries-on-connection-error"},
			},
			{
				input:       "pushsrv07.server-web_push-prod_07.retries",
				measurement: "cassandra",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "server-web_push-prod_07", "name": "prod_07", "type": "retries"},
			},
			{
				input:       "pushdemo00.channel-push_gcm-demo_00.GcmChannelSender.com_idamob_tinkoff_android_pro_1020.nonZeroNewMessageAvailableCount.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_gcm-demo_00", "name": "demo_00", "type": "count", "midname": "GcmChannelSender.com_idamob_tinkoff_android_pro_1020.nonZeroNewMessageAvailableCount"},
			},
			{
				input:       "pushsrv08.connector-gate-sb_06.deviceStatusOutPacketFailSafeTransmitter.meter.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv08", "component": "connector-gate-sb_06", "name": "sb_06", "type": "count", "midname": "deviceStatusOutPacketFailSafeTransmitter.meter"},
			},
			{
				input:       "pushdemo00.connector-http-brooma_demo_06.ConnectorOutMessageDlvStatusCounter.rejected.m1",
				measurement: "counter.status",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-http-brooma_demo_06", "name": "brooma_demo_06", "type": "m1", "midname": "ConnectorOutMessageDlvStatusCounter", "status": "rejected"},
			},
			{
				input:       "pushdemo00.channel-push_apnshttp-demo_00.StoredQueue.Transmitter.CmiPushMessageDlvEvent.connector-http-tcsbankdemo_00.size",
				measurement: "storedqueue.transmitter",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "name": "demo_00", "type": "size", "midname": "StoredQueue.Transmitter.CmiPushMessageDlvEvent.connector-http-tcsbankdemo_00"},
			},
			{
				input:       "pushdemo00.channel-push_apnshttp-demo_00.StoredQueue.655.size",
				measurement: "storedqueue",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "name": "demo_00", "type": "size", "midname": "StoredQueue.655"},
			},
			{
				input:       "pushdemo00.connector-gate-sovcombankdemo_00.PlatformPacketListener.errorPacketCount",
				measurement: "value",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sovcombankdemo_00", "name": "sovcombankdemo_00", "type": "errorPacketCount", "midname": "PlatformPacketListener"},
			},
			{
				input:       "pushdemo00.connector-gate-bcsdemo_00.PlatformPacketListener.outMessageTimer.p95",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-bcsdemo_00", "name": "bcsdemo_00", "type": "p95", "midname": "PlatformPacketListener.outMessageTimer"},
			},
			{
				input:       "pushdemo00.connector-gate-sbdemo_04.PlatformPacketListener.latestReceivedPacketTimestamp",
				measurement: "timestamp",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-sbdemo_04", "name": "sbdemo_04", "type": "latestReceivedPacketTimestamp", "midname": "PlatformPacketListener"},
			},
			{
				input:       "pushdemo00.channel-push_apnshttp-demo_00.ApnsHttpChannelSender.ru_sberbank_onlineiphone_beta_42.certificateValidDayCount",
				measurement: "value",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_apnshttp-demo_00", "name": "demo_00", "type": "certificateValidDayCount", "midname": "ApnsHttpChannelSender.ru_sberbank_onlineiphone_beta_42"},
			},
			{
				input:       "pushsrv05.channel-push_apnshttp-prod_05.ApnsHttpChannelSender.com_brooma_threads-beta_377.nonZeroNewMessageAvailableCount.m1",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv05", "component": "channel-push_apnshttp-prod_05", "name": "prod_05", "type": "m1", "midname": "ApnsHttpChannelSender.com_brooma_threads-beta_377.nonZeroNewMessageAvailableCount"},
			},
			{
				input:       "pushdemo00.connector-gate-openbankdemo_00.ClientMessageDatabaseProcessor.queueSize",
				measurement: "queue",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-openbankdemo_00", "name": "openbankdemo_00", "type": "queueSize", "midname": "ClientMessageDatabaseProcessor"},
			},
			{
				input:       "pushsrv03.channel-sms_hpx-prod_01.HpxChannelProcessor-psbank.processTimer.m1",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushsrv03", "component": "channel-sms_hpx-prod_01", "name": "prod_01", "type": "m1", "midname": "HpxChannelProcessor-psbank.processTimer"},
			},
			{
				input:       "pushsrv00.server-web_push-prod_00.Transmitter.EnricherRequestCompletableFuture.sb.processTimer.p99",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "server-web_push-prod_00", "name": "prod_00", "type": "p99", "midname": "Transmitter.EnricherRequestCompletableFuture.sb.processTimer"},
			},
			{
				input:       "pushdemo00.connector-gate-vtb24demo_02.Transmitter.EnricherRequestCompletableFuture.vtb24.enqueueTimer.p999",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-gate-vtb24demo_02", "name": "vtb24demo_02", "type": "p999", "midname": "Transmitter.EnricherRequestCompletableFuture.vtb24.enqueueTimer"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.SmsCanceled.Sb.Base.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "name": "demo_00", "type": "count", "midname": "SmsCanceled.Sb.Base"},
			},
			{
				input:       "pushsrv03.channel-sms_hpx-prod_01.HpxChannelProcessor-psbank.queueSize",
				measurement: "queue",
				tags:        map[string]string{"hostname": "pushsrv03", "component": "channel-sms_hpx-prod_01", "name": "prod_01", "type": "queueSize", "midname": "HpxChannelProcessor-psbank"},
			},
			{
				input:       "pushsrv03.connector-http-brooma_13.In_message_confirm.cnrId-131.acnId-17.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv03", "component": "connector-http-brooma_13", "name": "brooma_13", "type": "count", "midname": "In_message_confirm.cnrId-131.acnId-17"},
			},
			{
				input:       "pushdemo00.server-web_push-demo_00.EnricherRequestRemoteServiceClient.skippedRequestMeter.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00", "name": "demo_00", "type": "count", "midname": "EnricherRequestRemoteServiceClient.skippedRequestMeter"},
			},
			{
				input:       "pushdemo00.server-web_push-demo_00.EnricherRequestRemoteServiceClient.acnId-10.requestTimer.m1",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00", "name": "demo_00", "type": "m1", "midname": "EnricherRequestRemoteServiceClient.acnId-10.requestTimer"},
			},
			{
				input:       "pushsrv01.channel-sms_hpx-prod_00.SmsUndelivered.Sovcombank.Rosban.m1",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv01", "component": "channel-sms_hpx-prod_00", "name": "prod_00", "type": "m1", "midname": "SmsUndelivered.Sovcombank.Rosban"},
			},
			{
				input:       "pushsrv03.channel-sms_hpx-prod_01.HpxSentOutMessageMissedCache.cacheSize",
				measurement: "cache",
				tags:        map[string]string{"hostname": "pushsrv03", "component": "channel-sms_hpx-prod_01", "name": "prod_01", "type": "cacheSize", "midname": "HpxSentOutMessageMissedCache"},
			},
			{
				input:       "pushsrv06.channel-sms_hpx-prod_02.StoredQueue.HpxChannel-vtb24-queue.size",
				measurement: "storedqueue",
				tags:        map[string]string{"hostname": "pushsrv06", "component": "channel-sms_hpx-prod_02", "name": "prod_02", "type": "size", "midname": "StoredQueue.HpxChannel-vtb24-queue"},
			},
			{
				input:       "pushsrv05.channel-push_gcmxmpp-prod_05.StoredQueue.Transmitter.CmiPushMessageDlvEvent.connector-http-brooma_01.size",
				measurement: "storedqueue.transmitter",
				tags:        map[string]string{"hostname": "pushsrv05", "component": "channel-push_gcmxmpp-prod_05", "name": "prod_05", "type": "size", "midname": "StoredQueue.Transmitter.CmiPushMessageDlvEvent.connector-http-brooma_01"},
			},
			{
				input:       "pushsrv00.connector-http-brooma_12.SmsMessageTransmitterPool.errorCount.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "connector-http-brooma_12", "name": "brooma_12", "type": "count", "midname": "SmsMessageTransmitterPool.errorCount"},
			},
			{
				input:       "pushsrv00.connector-http-brooma_12.PushMessageTransmitterPool.aliveTransmittersCount",
				measurement: "pool",
				tags:        map[string]string{"hostname": "pushsrv00", "component": "connector-http-brooma_12", "name": "brooma_12", "type": "aliveTransmittersCount", "midname": "PushMessageTransmitterPool"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.Transmitter.CmiSmsMessageDlvEvent.connector-http-tcsbank_02.enqueueTimer.mean",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "name": "demo_00", "type": "mean", "midname": "Transmitter.CmiSmsMessageDlvEvent.connector-http-tcsbank_02.enqueueTimer"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.Transmitter.CmiSmsMessageDlvEvent.connector-db-test_demo.enqueueTimer.mean",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "name": "demo_00", "type": "mean", "midname": "Transmitter.CmiSmsMessageDlvEvent.connector-db-test_demo.enqueueTimer"},
			},
			{
				input:       "pushdemo00.server-web_push-demo_00.Transmitter.CmiPushResendEvent.connector-http-broomademo_08.consumerAvailable",
				measurement: "state",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "server-web_push-demo_00", "name": "demo_00", "type": "consumerAvailable", "midname": "Transmitter.CmiPushResendEvent.connector-http-broomademo_08"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.HpxOutMessageClientIdCache.idCacheSize",
				measurement: "idcache",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "name": "demo_00", "type": "idCacheSize", "midname": "HpxOutMessageClientIdCache"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.DataSourceFactoryBean.numActive",
				measurement: "bean",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "name": "demo_00", "type": "numActive", "midname": "DataSourceFactoryBean"},
			},
			{
				input:       "pushdemo00.channel-sms_hpx-demo_00.StoredMap.hpxOutMessageCache.size",
				measurement: "storedmap",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-sms_hpx-demo_00", "name": "demo_00", "type": "size", "midname": "StoredMap.hpxOutMessageCache"},
			},
			{
				input:       "pushsrv07.connector-httpxml_pushmessage-raiff_01.open-files.descriptors",
				measurement: "value",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "connector-httpxml_pushmessage-raiff_01", "name": "raiff_01", "type": "descriptors", "midname": "open-files"},
			},
			{
				input:       "pushsrv07.connector-httpxml_pushmessage-raiff_01.bytes-sent.m1",
				measurement: "cassandra.counter",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "connector-httpxml_pushmessage-raiff_01", "name": "raiff_01", "type": "m1", "midname": "bytes-sent"},
			},
			{
				input:       "pushsrv04.connector-gate-sb_05.gateProcessingTimer.PUSH_DEVICE_STATUS_UPDATE.p999",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushsrv04", "component": "connector-gate-sb_05", "name": "sb_05", "type": "p999", "midname": "gateProcessingTimer.PUSH_DEVICE_STATUS_UPDATE"},
			},
			{
				input:       "pushsrv07.connector-gate-sb_01.PushServerPacketProcessor.outQueueSize",
				measurement: "value",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "connector-gate-sb_01", "name": "sb_01", "type": "outQueueSize", "midname": "PushServerPacketProcessor"},
			},
			{
				input:       "pushsrv07.connector-gate-sb_01.TrafficMonitoringService.statusMeter.push_upstream_delivered.count",
				measurement: "counter.status",
				tags:        map[string]string{"hostname": "pushsrv07", "component": "connector-gate-sb_01", "name": "sb_01", "type": "count", "midname": "TrafficMonitoringService.statusMeter", "status": "push_upstream_delivered"},
			},
			{
				input:       "pushdemo00.channel-push_websocket-demo_00.WebSocketChannelSender.web_threads_chat_minbank_680.notificationsPushedWithError",
				measurement: "value",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "channel-push_websocket-demo_00", "name": "demo_00", "type": "notificationsPushedWithError", "midname": "WebSocketChannelSender.web_threads_chat_minbank_680"},
			},
			{
				input:       "pushsrv05.channel-push_gcmxmpp-prod_05.GcmXmppChannelSender.ru_domopult_mitino_android_467.notificationsEnqueued",
				measurement: "value",
				tags:        map[string]string{"hostname": "pushsrv05", "component": "channel-push_gcmxmpp-prod_05", "name": "prod_05", "type": "notificationsEnqueued", "midname": "GcmXmppChannelSender.ru_domopult_mitino_android_467"},
			},
			{
				input:       "pushsrv06.server-history_handler-prod_00.MessageHistoryMetricInfo.pushServerId-3.enqueued.count",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushsrv06", "component": "server-history_handler-prod_00", "name": "prod_00", "type": "count", "midname": "MessageHistoryMetricInfo.pushServerId-3.enqueued"},
			},
			{
				input:       "pushdemo00.connector-advisa-demo_00.ClientOutMessageDlvTimeCounter.enqueuedToDelivered.p999",
				measurement: "timer",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-advisa-demo_00", "name": "demo_00", "type": "p999", "midname": "ClientOutMessageDlvTimeCounter.enqueuedToDelivered"},
			},
			{
				input:       "pushdemo00.connector-db-test_demo.CassandraStorageWsPostponeMessageDataAccessor.successCount.m1",
				measurement: "counter",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-db-test_demo", "name": "test_demo", "type": "m1", "midname": "CassandraStorageWsPostponeMessageDataAccessor.successCount"},
			},
			{
				input:       "pushdemo00.connector-http-tcsbankdemo_01.ClientOutMessageDlvStatusCounter.failed.count",
				measurement: "counter.status",
				tags:        map[string]string{"hostname": "pushdemo00", "component": "connector-http-tcsbankdemo_01", "name": "tcsbankdemo_01", "type": "count", "midname": "ClientOutMessageDlvStatusCounter", "status": "failed"},
			},
			{
				input:       "pushsrv02.connector-gate-psbank_01.TrafficMonitoringService.statusMeter.push_upstream_delivered.m1",
				measurement: "counter.status",
				tags:        map[string]string{"hostname": "pushsrv02", "component": "connector-gate-psbank_01", "name": "psbank_01", "type": "m1", "midname": "TrafficMonitoringService.statusMeter", "status": "push_upstream_delivered"},
			},
		},
		"advisa": {
			{
				input:       "avitst00.advisa-database-monitor-test_00.DatabasePingService.avgBusyConnections.count",
				measurement: "counter",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "count", "midname": "DatabasePingService.avgBusyConnections"},
			},
			{
				input:       "avitst00.advisa-database-monitor-test_00.DataSourceFactoryBean.numIdle",
				measurement: "bean",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "numIdle", "midname": "DataSourceFactoryBean"},
			},
			{
				input:       "avitst00.advisa-database-monitor-test_00.non-heap.max",
				measurement: "memory",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "max", "midname": "non-heap"},
			},
			{
				input:       "avitst00.advisa-database-monitor-test_00.pools.PS-Survivor-Space.used-after-gc",
				measurement: "memory",
				tags:        map[string]string{"component": "advisa-database-monitor-test_00", "hostname": "avitst00", "type": "used-after-gc", "midname": "pools.PS-Survivor-Space"},
			},
			{
				input:       "avitst00.connector-terminal-test_00.heap.used",
				measurement: "memory",
				tags:        map[string]string{"component": "connector-terminal-test_00", "hostname": "avitst00", "type": "used", "midname": "heap"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.advisaOperationProcessor.retryQueueSize",
				measurement: "queue",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00", "type": "retryQueueSize", "midname": "advisaOperationProcessor"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.coasConnectorOperationDlvEventConsumer.enqueueTimer.p999",
				measurement: "timer",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00", "type": "p999", "midname": "coasConnectorOperationDlvEventConsumer.enqueueTimer"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.Transmitter.CoasSecurityTokenUpdate.connector-advisa-test_00.enqueueTimer.p99",
				measurement: "transmitter.timer",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00", "type": "p99", "midname": "Transmitter.CoasSecurityTokenUpdate.connector-advisa-test_00.enqueueTimer"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.CoasSubscriptionStatusEventTransmitterPool.aliveTransmittersCount",
				measurement: "pool",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00", "type": "aliveTransmittersCount", "midname": "CoasSubscriptionStatusEventTransmitterPool"},
			},
			{
				input:       "avitst00.connector-terminal_location_aggregator-test_00.StoredQueue.terminalLocationsAggregateMultiPointProcessorRetry.size",
				measurement: "storedqueue",
				tags:        map[string]string{"component": "connector-terminal_location_aggregator-test_00", "hostname": "avitst00", "type": "size", "midname": "StoredQueue.terminalLocationsAggregateMultiPointProcessorRetry"},
			},
			{
				input:       "avitst00.server-web_advisa-test_00.uptime",
				measurement: "uptime",
				tags:        map[string]string{"component": "server-web_advisa-test_00", "hostname": "avitst00"},
			},
			{
				input:       "avisrv00.connector-operations_smsparse-prod_00.Transmitter.CoasCardTransaction.server-web_advisa-prod_01.enqueueTimer.p95",
				measurement: "transmitter.timer",
				tags:        map[string]string{"component": "connector-operations_smsparse-prod_00", "hostname": "avisrv00", "type": "p95", "midname": "Transmitter.CoasCardTransaction.server-web_advisa-prod_01.enqueueTimer"},
			},
			{
				input:       "avisrv01.server-web_advisa-prod_01.Transmitter.CoasSecurityTokenUpdate.connector-advisa-prod_00.errorCount.count",
				measurement: "transmitter.counter",
				tags:        map[string]string{"component": "server-web_advisa-prod_01", "hostname": "avisrv01", "type": "count", "midname": "Transmitter.CoasSecurityTokenUpdate.connector-advisa-prod_00.errorCount"},
			},
			{
				input:       "avisrv00.connector-operations_smsparse-prod_00.Transmitter.AdariAdvisaOutMessageDlvEvent.avirouter1.processTimer.p95",
				measurement: "transmitter.timer",
				tags:        map[string]string{"component": "connector-operations_smsparse-prod_00", "hostname": "avisrv00", "type": "p95", "midname": "Transmitter.AdariAdvisaOutMessageDlvEvent.avirouter1.processTimer"},
			},
			{
				input:       "avisrv00.connector-registration-prod_00.Transmitter.CoasAdvisaRegistrationConfirmationResponse.server-web_advisa-prod_00.enqueueTimer.p999",
				measurement: "transmitter.timer",
				tags:        map[string]string{"component": "connector-registration-prod_00", "hostname": "avisrv00", "type": "p999", "midname": "Transmitter.CoasAdvisaRegistrationConfirmationResponse.server-web_advisa-prod_00.enqueueTimer"},
			},
			{
				input:       "avisrv01.connector-operations_smsparse-prod_01.StoredQueue.Transmitter.AdariAdvisaOutMessageDlvEvent.avirouter0.size",
				measurement: "storedqueue.transmitter",
				tags:        map[string]string{"component": "connector-operations_smsparse-prod_01", "hostname": "avisrv01", "type": "size", "midname": "StoredQueue.Transmitter.AdariAdvisaOutMessageDlvEvent.avirouter0"},
			},
			{
				input:       "avisrv00.connector-operations_smsparse-prod_00.Transmitter.AdariAdvisaOutMessageDlvEvent.avirouter0.queueSize",
				measurement: "transmitter",
				tags:        map[string]string{"component": "connector-operations_smsparse-prod_00", "hostname": "avisrv00", "type": "queueSize", "midname": "Transmitter.AdariAdvisaOutMessageDlvEvent.avirouter0"},
			},
			{
				input:       "avitst00.server-web_advisa-test_01.StoredQueue.operationReceivedProcessorNGRetry.size",
				measurement: "storedqueue",
				tags:        map[string]string{"component": "server-web_advisa-test_01", "hostname": "avitst00", "type": "size", "midname": "StoredQueue.operationReceivedProcessorNGRetry"},
			},
			{
				input:       "avitst00.connector-registration-test_00.Transmitter.CoasAdvisaRegistrationConfirmationResponse.server-web_advisa-test_00.queueSize",
				measurement: "transmitter",
				tags:        map[string]string{"component": "connector-registration-test_00", "hostname": "avitst00", "type": "queueSize", "midname": "Transmitter.CoasAdvisaRegistrationConfirmationResponse.server-web_advisa-test_00"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.StoredQueue.Transmitter.CoasAdvisaRegistrationConfirmationResponse.server-web_advisa-prod_01.size",
				measurement: "storedqueue.transmitter",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "size", "midname": "StoredQueue.Transmitter.CoasAdvisaRegistrationConfirmationResponse.server-web_advisa-prod_01"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.StoredMap.MerchantDTO.size",
				measurement: "storedmap",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "size", "midname": "StoredMap.MerchantDTO"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.SystemJournalService.writeFirstSyncIfNeededDuration.mean",
				measurement: "timer",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "mean", "midname": "SystemJournalService.writeFirstSyncIfNeededDuration"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.TerminalFullTextIndexingProcessor.notIndexedCount",
				measurement: "value",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "notIndexedCount", "midname": "TerminalFullTextIndexingProcessor"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.coasAdvisaPushServerDeviceStatusEventConsumer.avgSpeed",
				measurement: "speed",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "avgSpeed", "midname": "coasAdvisaPushServerDeviceStatusEventConsumer"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.coasAdvisaSubscriptionStatusEventConsumerImpl.queueSize",
				measurement: "queue",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "queueSize", "midname": "coasAdvisaSubscriptionStatusEventConsumerImpl"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.PushMessageIdCache.idCacheSize",
				measurement: "idcache",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "idCacheSize", "midname": "PushMessageIdCache"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.HpxSentOutMessageMissedCache.missCount",
				measurement: "cache",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "missCount", "midname": "HpxSentOutMessageMissedCache"},
			},
			{
				input:       "avisrv00.connector-registration-rbr_00.Transmitter.CoasSubscriptionStatusEvent.server-web_advisa-prod_01.processedCount.count",
				measurement: "transmitter.counter",
				tags:        map[string]string{"component": "connector-registration-rbr_00", "hostname": "avisrv00", "type": "count", "midname": "Transmitter.CoasSubscriptionStatusEvent.server-web_advisa-prod_01.processedCount"},
			},
			{
				input:       "avitst00.connector-subscription-test_00.PS-MarkSweep.count",
				measurement: "gc",
				tags:        map[string]string{"component": "connector-subscription-test_00", "hostname": "avitst00", "type": "count", "midname": "PS-MarkSweep"},
			},
			{
				input:       "avitst00.server-account_grouping-test_00.open-files.descriptors",
				measurement: "value",
				tags:        map[string]string{"component": "server-account_grouping-test_00", "hostname": "avitst00", "type": "descriptors", "midname": "open-files"},
			},
			{
				input:       "avitst00.connector-operations_smsparse-test_00.CoasConnectorBankAdvisaDlvEventConsumerImpl.timer.p75",
				measurement: "timer",
				tags:        map[string]string{"component": "connector-operations_smsparse-test_00", "hostname": "avitst00", "type": "p75", "midname": "CoasConnectorBankAdvisaDlvEventConsumerImpl.timer"},
			},
			{
				input:       "avisrv01.connector-terminal-prod_00.TerminalSimilarIndexService.successRate",
				measurement: "index",
				tags:        map[string]string{"component": "connector-terminal-prod_00", "hostname": "avisrv01", "type": "successRate", "midname": "TerminalSimilarIndexService"},
			},
			{
				input:       "avisrv01.server-web_admin-test_00.UploadTerminalsService.avgOverallSpeedMeter.count",
				measurement: "counter",
				tags:        map[string]string{"component": "server-web_admin-test_00", "hostname": "avisrv01", "type": "count", "midname": "UploadTerminalsService.avgOverallSpeedMeter"},
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
				t.Fatalf("name parse failed. expected %v got %v, input %v", test.measurement, measurement, test.input)
			}
			if len(tags) != len(test.tags) {
				t.Fatalf("unexpected number of tags.  expected %v, got %v, input %v", test.tags, tags, test.input)
			}
			for k, v := range test.tags {
				if tags[k] != v {
					t.Fatalf("unexpected tag value for tags[%s].  expected %q, got %q, input %v", k, v, tags[k], test.input)
				}
			}
		}
	}
}
