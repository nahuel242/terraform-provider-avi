package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// EventLog event log
// swagger:model EventLog
type EventLog struct {

	//  Enum options - EVENT_CONTEXT_SYSTEM, EVENT_CONTEXT_CONFIG, EVENT_CONTEXT_APP, EVENT_CONTEXT_ALL.
	Context *string `json:"context,omitempty"`

	// Summary of event details.
	DetailsSummary *string `json:"details_summary,omitempty"`

	// Event Description for each Event  in the table view.
	EventDescription *string `json:"event_description,omitempty"`

	// Placeholder for description of property event_details of obj type EventLog field type str  type object
	EventDetails *EventDetails `json:"event_details,omitempty"`

	//  Enum options - VINFRA_DISC_DC, VINFRA_DISC_HOST, VINFRA_DISC_CLUSTER, VINFRA_DISC_VM, VINFRA_DISC_NW, MGMT_NW_NAME_CHANGED, DISCOVERY_DATACENTER_DEL, VM_ADDED, VM_REMOVED, VINFRA_DISC_COMPLETE, VCENTER_ADDRESS_ERROR, SE_GROUP_CLUSTER_DEL, SE_GROUP_MGMT_NW_DEL, MGMT_NW_DEL, VCENTER_BAD_CREDENTIALS, ESX_HOST_UNREACHABLE, SERVER_DELETED, SE_GROUP_HOST_DEL, VINFRA_DISC_FAILURE, ESX_HOST_POWERED_DOWN, VCENTER_VERSION_NOT_SUPPORTED, VCENTER_CONNECTIVITY_FAIL, VCENTER_CONNECTIVITY_SUCCESS, VCENTER_ACCESS_SLOW, VCENTER_USER_ROLE_CHANGE, VCENTER_NETWQRK_OBJECT_LIMIT_REACHED, SE_FATAL_ERROR, SE_HEARTBEAT_FAILURE, SE_MARKED_DOWN, SE_VM_DELETED, SE_VM_PURGED, SE_UP, SE_POWERED_DOWN, SE_REBOOTED, SE_HEALTH_CHECK_FAIL, SE_EXTERNAL_HM_RESTART, SE_DOWN, SE_VERSION_CHECK_FAILED, SE_UPGRADING, SE_ENABLE, SE_MIGRATE, CREATING_SE, CREATED_SE, CREATE_SE_FAIL, CREATE_SE_TIMEOUT, DELETING_SE, DELETED_SE, DELETE_SE_FAIL, ADD_NW_SE, DEL_NW_SE, VS_ADD_SE_INT, VS_REMOVED_SE_INT, VS_ADD_SE, VS_REMOVED_SE, ADD_NW_FAIL, RM_DEL_NETWORK_FAIL, REBOOT_SE, MODIFY_NW, MODIFY_NW_FAIL, VS_SE_BOOTUP_FAIL, VS_SE_IP_FAIL, NO_HOST_AVAIL, VS_SWITCHOVER, VS_SWITCHOVER_FAIL, ADD_VIP_VNIC, DEL_VIP_VNIC, VS_FSM_INACTIVE, VS_FSM_AWAITING_SE_ASSIGNMENT, VS_FSM_ACTIVE, VS_FSM_ACTIVE_AWAITING_SE_TRANSITION, VS_FSM_DISABLED, NEW_PROBABLE_SRVR, VS_SCALEOUT_DONE, VS_SCALEOUT_DONE_AWAITING_MORE_SE, VS_SCALEOUT_ERR, VS_SCALEIN_DONE, VS_SCALEIN_DONE_AWAITING_MORE_SE, VS_SCALEIN_ERR, VS_MIGRATE_SCALEOUT_DONE, VS_MIGRATE_SCALEOUT_ERROR, VS_MIGRATE_SCALEIN_DONE, VS_MIGRATE_SCALEIN_ERROR, VS_MIGRATE_DONE, VS_FSM_UNEXPECTED_EVENT, VS_RPC_TO_RESMGR_FAILED_EVENT, VS_RPC_TO_SE_FAILED_EVENT, VS_RPC_FAILED_EVENT, VS_SCALEOUT_COMPLETE, VS_SCALEIN_COMPLETE, VS_MIGRATE_STARTED, VS_MIGRATE_COMPLETE, VS_SCALEOUT_FAILED, VS_SCALEIN_FAILED, VS_MIGRATE_FAILED, VS_AWAITING_SE, VS_INITIAL_PLACEMENT_FAILED, VS_FSM_ACTIVE_AWAITING_SCALEOUT_READY, UPGRADE_ALL_SE_START, UPGRADE_ALL_SE_DONE, UPGRADE_ALL_SE_NOT_NEEDED, UPGRADE_SE_START, UPGRADE_SE_DONE, UPGRADE_SE_NOT_NEEDED, UPGRADE_SE_SUSPENDED, UPGRADE_SE_VS_SCALEOUT, UPGRADE_SE_VS_SCALEIN, UPGRADE_SE_VS_MIGRATE, UPGRADE_SE_VS_DISRUPTED, REBALANCE_VS_SCALEOUT, REBALANCE_VS_SCALEIN, REBALANCE_VS_MIGRATE, DISABLE_SE_VS_MIGRATE, ROLLBACK_ALL_SE_START, ROLLBACK_ALL_SE_DONE, MIGRATE_SE_STARTED, MIGRATE_SE_RESTARTED, MIGRATE_SE_FINISHED, MIGRATE_SE_FAILED, MIGRATE_SE_VS_MIGRATE_STARTED, MIGRATE_SE_VS_MIGRATE_FINISHED, MIGRATE_SE_VS_MIGRATE_FAILED, VIP_SCALEOUT, VIP_SCALEOUT_FAILED, VIP_SCALEIN, VIP_SCALEIN_FAILED, SE_HM_EVENT_SHM_DOWN, SE_HM_EVENT_SHM_UP, SERVER_DOWN, SERVER_UP, POOL_DOWN, POOL_UP, VS_DOWN, VS_UP, SE_SERVER_DELETED, SE_SERVER_DISABLED, SE_POOL_DELETED, SE_SERVER_APP_CHANGED, VS_CONN_LIMIT, VS_THROUGHPUT_LIMIT, CONN_DROP_MAX_SYN_TBL, CONN_DROP_MAX_FLOW_TBL, CONN_DROP_MAX_PERSIST_TBL, CONN_DROP_POOL_LB_FAILURE, CONN_DROP_NO_CONN_MEM, CONN_DROP_NO_PKT_BUFF, PKT_DROP_NO_PKT_BUFF, PKT_BUFF_ALLOC_FAIL, CACHE_OBJ_ALLOC_FAIL, SE_CPU_HIGH, SE_MEM_HIGH, SE_PKT_BUFF_HIGH, SE_PERSIST_TBL_HIGH, SE_CONN_MEM_HIGH, SE_DISK_HIGH, SE_FLOW_TBL_HIGH, SE_SYN_TBL_HIGH, SE_DP_HB_FAILED, SE_VNIC_DHCP_IP_ALLOC_FAILURE, SE_VNIC_DUPLICATE_IP, SE_SYN_CACHE_USAGE_HIGH, VS_SE_HA_ACTIVE, VS_SE_HA_COMPROMISED, POOL_SE_HA_ACTIVE, POOL_SE_HA_COMPROMISED, SERVER_DOWN_HA_COMPROMISED, SERVER_UP_HA_ACTIVE, SE_VNIC_IP_ADDED, SE_VNIC_IP_REMOVED, GS_MEMBER_DOWN, GS_MEMBER_UP, GS_GROUP_DOWN, GS_GROUP_UP, GS_DOWN, GS_UP, VIP_DOWN, VIP_UP, SE_GEO_DB_FAILURE, VS_GEO_DB_FAILURE, SE_GEO_DB_SUCCESS, VS_GEO_DB_SUCCESS, SE_EV_SERVER_DOWN, SE_EV_SERVER_UP, SE_EV_POOL_DOWN, SE_EV_POOL_UP, SE_EV_VS_DOWN, SE_EV_VS_UP, SE_HM_EVENT_GHM_DOWN, SE_HM_EVENT_GHM_UP, SE_EV_GS_GROUP_DELETED, SE_EV_GS_MEMBER_DOWN, SE_EV_GS_MEMBER_UP, SE_EV_GS_GROUP_DOWN, SE_EV_GS_GROUP_UP, SE_EV_GS_DOWN, SE_EV_GS_UP, SE_IP6_DAD_FAILED, CONFIG_CREATE, CONFIG_UPDATE, CONFIG_DELETE, USER_LOGIN, USER_LOGOUT, CONFIG_ACTION, CONFIG_INTERNAL_CREATE, CONFIG_INTERNAL_UPDATE, USER_PASSWORD_CHANGE_REQUEST, USER_AUTHORIZED_BY_RULE, USER_NOT_AUTHORIZED_BY_ANY_RULE, CONFIG_SE_GRP_FLAVOR_UPDATE, SSL_CERT_EXPIRE, SSL_KEY_EXPORTED, SSL_CERT_RENEW, SSL_CERT_RENEW_FAILED, CONTROLLER_NODE_JOINED, CONTROLLER_NODE_LEFT, CONTROLLER_SERVICE_FAILURE, CONTROLLER_LEADER_FAILOVER, CONTROLLER_WARM_REBOOT, CONTROLLER_SERVICE_RESTORED, CONTROLLER_SERVICE_CRITICAL_FAILURE, CONTROLLER_NODE_SHUTDOWN, CONTROLLER_NODE_STARTED, CLUSTER_CONFIG_FAILED, SYSTEM_UPGRADE_STARTED, SYSTEM_UPGRADE_COMPLETE, SYSTEM_UPGRADE_ABORTED, SYSTEM_ROLLBACK_STARTED, SYSTEM_ROLLBACK_COMPLETE, SYSTEM_ROLLBACK_ABORTED, CONTROLLER_NODE_DB_REPLICATION_FAILED, CONTROLLER_PROCESS_STOPPED, CONTROLLER_MEMORY_BALANCER_DISABLED, METRIC_THRESHOLD_UP_VIOLATION, LICENSE_EXPIRY, ANOMALY, LICENSE_ADDITION_NOTIF, LICENSE_REMOVAL_NOTIF, METRICS_DB_DISK_FULL, OPENSTACK_ACCESS_FAILURE, OPENSTACK_ACCESS_SUCCESS, OPENSTACK_IMAGE_UPLOAD_FAILURE, OPENSTACK_IMAGE_UPLOAD_SUCCESS, OPENSTACK_SE_VM_CREATED, OPENSTACK_SE_VM_DELETED, OPENSTACK_SE_VM_DELETION_DETECTED, OPENSTACK_VNIC_ADDED, OPENSTACK_VNIC_REMOVED, OPENSTACK_IP_DETACHED, OPENSTACK_IP_ATTACHED, OPENSTACK_SE_CREATION_FAILURE, OPENSTACK_SE_DELETION_FAILURE, OPENSTACK_VNIC_ADDITION_FAILURE, OPENSTACK_VNIC_DELETION_FAILURE, OPENSTACK_IP_DETACH_FAILURE, OPENSTACK_IP_ATTACH_FAILURE, OPENSTACK_LBPROV_AUDIT_FAILURE, OPENSTACK_LBPROV_AUDIT_SUCCESS, OPENSTACK_LBPLUGIN_OP_FAILURE, OPENSTACK_LBPLUGIN_OP_SUCCESS, OPENSTACK_SYNC_SERVICES_SUCCESS, OPENSTACK_SYNC_SERVICES_FAILURE, OPENSTACK_TENANTS_DELETED, AWS_ACCESS_FAILURE, AWS_ACCESS_SUCCESS, AWS_IMAGE_UPLOAD_FAILURE, AWS_IMAGE_UPLOAD_SUCCESS, AWS_SNS_ACCESS_FAILURE, AWS_SNS_ACCESS_SUCCESS, AWS_SQS_ACCESS_FAILURE, AWS_SQS_ACCESS_SUCCESS, AWS_ASG_PUT_NOTIFICATION_CONFIGURATION_FAILURE, AWS_ASG_PUT_NOTIFICATION_CONFIGURATION_SUCCESS, AWS_ASG_DELETE_NOTIFICATION_CONFIGURATION_FAILURE, AWS_ASG_DELETE_NOTIFICATION_CONFIGURATION_SUCCESS, AWS_ASG_NOTIFICATION_PROCESSING_FAILURE, AWS_ASG_NOTIFICATION_PROCESSING_SUCCESS, AWS_ASG_NOTIFICATION_INSTANCE_ADDED, AWS_ASG_NOTIFICATION_INSTANCE_REMOVED, AWS_ASG_ACCESS_FAILURE, AWS_ASG_ACCESS_SUCCESS, AWS_ASG_NOTIFICATION_INSTANCE_LAUNCH_ERROR, AWS_ASG_NOTIFICATION_INSTANCE_TERMINATE_ERROR, AWS_ASG_NOTIFICATION_AUTOSCALE_GROUP_DELETED, CLOUDSTACK_ACCESS_FAILURE, CLOUDSTACK_ACCESS_SUCCESS, CLOUDSTACK_IMAGE_UPLOAD_FAILURE, CLOUDSTACK_IMAGE_UPLOAD_SUCCESS, DOCKER_UCP_ACCESS_SUCCESS, DOCKER_UCP_ACCESS_FAILURE, DOCKER_UCP_IMAGE_UPLOAD_FAILURE, DOCKER_UCP_IMAGE_UPLOAD_SUCCESS, DOCKER_UCP_IMAGE_UPLOAD_IN_PROGRESS, VCA_ACCESS_FAILURE, VCA_ACCESS_SUCCESS, VCA_IMAGE_UPLOAD_FAILURE, VCA_IMAGE_UPLOAD_SUCCESS, LS_ACCESS_FAILURE, LS_ACCESS_SUCCESS, LS_IMAGE_UPLOAD_FAILURE, LS_IMAGE_UPLOAD_SUCCESS, MESOS_ACCESS_SUCCESS, MESOS_ACCESS_FAILURE, MESOS_IMAGE_UPLOAD_FAILURE, MESOS_IMAGE_UPLOAD_SUCCESS, MESOS_IMAGE_UPLOAD_IN_PROGRESS, MESOS_CREATED_SE, MESOS_CREATE_SE_FAIL, MESOS_DELETED_SE, MESOS_DELETE_SE_FAIL, MESOS_STOPPED_SE, MESOS_STOP_SE_FAIL, MESOS_STARTED_SE, MESOS_START_SE_FAIL, MESOS_UPDATED_HOSTS, CC_SE_CREATED, CC_SE_CREATION_FAILURE, CC_SE_DELETED, CC_SE_DELETION_FAILURE, CC_SE_DELETION_DETECTED, CC_VNIC_ADDED, CC_VNIC_ADDITION_FAILURE, CC_VNIC_DELETED, CC_VNIC_DELETION_FAILURE, CC_IP_ATTACHED, CC_IP_ATTACH_FAILURE, CC_IP_DETACHED, CC_IP_DETACH_FAILURE, CC_SYNC_SERVICES_SUCCESS, CC_SYNC_SERVICES_FAILURE, CC_UPDATE_VIP_FAILURE, CC_DELETE_VIP_FAILURE, CC_CONFIG_FAILURE, CC_DECONFIG_FAILURE, CC_GENERIC_FAILURE, CC_CLUSTER_VIP_CONFIG_SUCCESS, CC_CLUSTER_VIP_CONFIG_FAILURE, CC_CLUSTER_VIP_DECONFIG_SUCCESS, CC_CLUSTER_VIP_DECONFIG_FAILURE, CC_MARATHON_SERVICE_PORT_OUTSIDE_VALID_RANGE, CC_MARATHON_SERVICE_PORT_ALREADY_IN_USE, CC_VIP_DNS_REGISTER_FAILURE, CC_TENANT_INIT_FAILURE, CC_HEALTH_FAILURE, CC_HEALTH_OK, CC_SE_STARTED, CC_SE_START_FAILURE, CC_SE_STOPPED, CC_SE_STOP_FAILURE, CC_VIP_PARK_INTF_SUCCESS, CC_VIP_PARK_INTF_FAILURE, CC_VIP_DNS_DEREGISTER_FAILURE, CC_VIP_DNS_VALIDATION_FAILURE, CC_VIP_DNS_REGISTER_SUCCESS, CC_VIP_DNS_DEREGISTER_SUCCESS, AWS_ROUTE53_ACCESS_FAILURE, AWS_ROUTE53_ACCESS_SUCCESS, VS_HEALTH_CHANGE, SE_HEALTH_CHANGE, POOL_HEALTH_CHANGE, SERVER_HEALTH_CHANGE, VS_HEALTH_DEGRADED, SE_HEALTH_DEGRADED, POOL_HEALTH_DEGRADED, SERVER_HEALTH_DEGRADED, DUPLICATE_SUBNETS, SUMMARIZED_SUBNETS, IP_POOL_ALMOST_EXHAUSTED, IP_POOL_EXHAUSTED, LICENSE_LIMIT_SERVERS, LICENSE_LIMIT_SE_VCPUS, LICENSE_LIMIT_THROUGHPUT, LICENSE_LIMIT_VS, LICENSE_LIMIT_HOSTS, LICENSE_LIMIT_SE_SOCKETS, LICENSE_EXPIRED, BURST_RESOURCE_CONSUMED, BURST_RESOURCE_EXPIRY_ALERT, APIC_BAD_CREDENTIALS, APIC_CREATE_LIFS, APIC_DELETE_LIFS, APIC_CREATE_LIF_CONTEXTS, APIC_DELETE_LIF_CONTEXTS, APIC_CREATE_CDEV, APIC_DELETE_CDEV, APIC_ATTACH_CIF_TO_LIF, APIC_DETACH_CIF_FROM_LIF, APIC_VS_PLACEMENT, APIC_BIND_VNIC_TO_NETWORK, APIC_CREATE_TENANT, APIC_DELETE_TENANT, APIC_CREATE_NETWORK, APIC_DELETE_NETWORK, APIC_NETWORK_VRF_CHANGED, APIC_VS_NETWORK_RESOLVE_ERROR, CONTAINER_CLOUD_ACCESS_SUCCESS, CONTAINER_CLOUD_ACCESS_FAILURE, CONTAINER_CLOUD_IMAGE_UPLOAD_FAILURE, CONTAINER_CLOUD_IMAGE_UPLOAD_SUCCESS, CONTAINER_CLOUD_IMAGE_UPLOAD_IN_PROGRESS, CONTAINER_CLOUD_CREATED_SE, CONTAINER_CLOUD_CREATE_SE_FAIL, CONTAINER_CLOUD_DELETED_SE, CONTAINER_CLOUD_DELETE_SE_FAIL, CONTAINER_CLOUD_STOPPED_SE, CONTAINER_CLOUD_STOP_SE_FAIL, CONTAINER_CLOUD_STARTED_SE, CONTAINER_CLOUD_START_SE_FAIL, CONTAINER_CLOUD_UPDATED_HOSTS, CONTAINER_CLOUD_SERVICE_SUCCESS, CONTAINER_CLOUD_SERVICE_FAILURE, CONTAINER_CLOUD_SERVICE_INCOMPLETE, AVG_UPTIME_CHANGE, DOS_ATTACK, SE_DOS_ATTACK, SERVER_AUTOSCALE_OUT, SERVER_AUTOSCALE_IN, SERVER_AUTOSCALE_OUT_COMPLETE, SERVER_AUTOSCALE_IN_COMPLETE, SERVER_AUTOSCALE_FAILED, SERVER_AUTOSCALE_IN_FAILED, SERVER_AUTOSCALE_OUT_FAILED, SE_GATEWAY_HEARTBEAT_FAILED, SE_GATEWAY_HEARTBEAT_SUCCESS, SE_VNIC_DOWN_EVENT, SE_VNIC_TX_QUEUE_STALL, SE_BGP_PEER_STATE_CHANGE, SE_LICENSED_BANDWIDTH_EXCEEDED, SERVER_AUTOSCALE_OUT_TRIGGERED, SERVER_AUTOSCALE_IN_TRIGGERED, POOL_AUTO_DEPLOYMENT_FAILED, POOL_AUTO_DEPLOYMENT_SUCCESS, SE_VNIC_UP_EVENT, POOL_AUTO_DEPLOYMENT_UPDATE, GSLB_SITE_OPER_STATUS, GSLB_DNS_STATUS, GSLB_SITE_EXCEPTION_STATUS, GSLB_GS_STATUS, SCHEDULER_ACTION_SUCCESS, SCHEDULER_ACTION_FAILURE, CONTROLLER_SCHEDULER_UNENCRYPTED_CONFIG_EXPORT, GCP_ACCESS_SUCCESS, GCP_ACCESS_FAIL, GCP_SE_DETECTED, GCP_API_FAIL, GCP_SUBNET_NOT_FOUND, GCP_SUBNET_ATTACH_FAIL, GCP_ROUTE_ADD_SUCCESS, GCP_ROUTE_DELETE_SUCCESS, GCP_ROUTE_ADD_FAIL, GCP_ROUTE_DELETE_FAIL, VIP_DNS_REGISTER_SUCCESS, VIP_DNS_REGISTER_FAILURE, VIP_DNS_DEREGISTER_SUCCESS, VIP_DNS_DEREGISTER_FAILURE, SYNC_DNS_RECORDS_SUCCESS, SYNC_DNS_RECORDS_FAILURE, FLUSH_DNS_RECORDS_SUCCESS, FLUSH_DNS_RECORDS_FAILURE, CC_HOST_SSH_FAILURE, CC_HOST_SSH_SUCCESS, AZURE_ACCESS_SUCCESS, AZURE_ACCESS_FAILURE, AZURE_ALB_UPDATE_FAILURE, AZURE_NIC_UPDATE_FAILURE, AZURE_ALB_UPDATE_SUCCESS, AZURE_NIC_UPDATE_SUCCESS, AZURE_NIC_DELETE_SUCCESS, AZURE_NIC_DELETE_FAILURE, AZURE_IMAGE_UPLOAD_FAILURE, AZURE_IMAGE_UPLOAD_SUCCESS, VS_FAULT, SE_SHM_MEM_HIGH, OCI_ACCESS_SUCCESS, OCI_ACCESS_FAILURE.
	// Required: true
	EventID *string `json:"event_id"`

	// Pages in which event should come up.
	EventPages []string `json:"event_pages,omitempty"`

	// Placeholder for description of property ignore_event_details_display of obj type EventLog field type str  type boolean
	IgnoreEventDetailsDisplay *bool `json:"ignore_event_details_display,omitempty"`

	//  Enum options - EVENT_INTERNAL, EVENT_EXTERNAL.
	Internal *string `json:"internal,omitempty"`

	// Placeholder for description of property is_security_event of obj type EventLog field type str  type boolean
	IsSecurityEvent *bool `json:"is_security_event,omitempty"`

	//  Enum options - UNKNOWN, VSMGR, SEMGR, RESMGR, VIMGR, METRICSMGR, CONFIG, SE_GENERAL, SE_FLOWTABLE, SE_HM, SE_POOL_PERSISTENCE, SE_POOL, VSERVER, CLOUD_CONNECTOR, CLUSTERMGR, HSMGR, NW_MGR, LICENSE_MGR, RES_MONITOR, STATEDBCACHE, STATEDBCACHEHA, APIC_AGENT, AUTOSCALE_MGR, GLB_MGR.
	// Required: true
	Module *string `json:"module"`

	// obj_name of EventLog.
	ObjName *string `json:"obj_name,omitempty"`

	ObjType *string `json:"obj_type,omitempty"`

	// Unique object identifier of obj.
	ObjUUID *string `json:"obj_uuid,omitempty"`

	// Reason code for generating the event. This would be added to the alert where it would say alert generated  on event with reason <reason code>. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_BAD_REQUEST, SYSERR_TEST1, SYSERR_TEST2, SYSERR_QUEUE_TRANSPORT_FAILURE, SYSERR_QUEUE_RETRY_TASK, SYSERR_DATASTORE_TRANSPORT_FAILURE, SYSERR_DATASTORE_UNKNOWN_FAILURE, SYSERR_DATASTORE_OBJECT_DOES_NOT_EXIST, SYSERR_DATASTORE_REFERENCE_DOES_NOT_EXIST, SYSERR_DATASTORE_DB_LOCKED, SYSERR_DATASTORE_LOCK_FAILURE, SYSERR_DATASTORE_TBL_NOT_EXIST, SYSERR_SVC_COMMON_OBJECT_NOT_IN_CACHED_VIEW, SYSERR_RPC_CANCELED_BY_CLIENT, SYSERR_RPC_TIMED_OUT, SYSERR_RPC_SEND_FAILED, SYSERR_RPC_CANCELED_BY_TRANSACTION_CLEANUP, SYSERR_NO_MULTICAST_RECEIVERS, SYSERR_RPC_FAILED, SYSERR_RPC_CONNECT_FAILED, SYSERR_CONTROLLER_NOT_READY, SYSERR_VERSION_MISMATCH, SYSERR_ALREADY_REGISTERED, SYSERR_INVALID_METHOD, SYSERR_DESERIALIZATION, SYSERR_SERIALIZATION, SYSERR_ENQUEUE, SYSERR_DEQUEUE, SYSERR_INVALID_READ_LEVEL, SYSERR_VS_INVALID_METHOD, SYSERR_VS_NOT_PRESENT, SYSERR_VS_INVALID_REQUEST, SYSERR_VS_NOT_ENOUGH_RESOURCES, SYSERR_VS_SE_NOT_AVAILABLE, SYSERR_VS_VNIC_FAILURE, SYSERR_VS_DELETE_WHILE_STILL_BEING_REFERRED, SYSERR_INVALID_HEALTH_MONITOR_TYPE, SYSERR_VS_SE_ASSIGNMENT_FAILED, SYSERR_VS_INVALID_OBJECT, SYSERR_VS_SERVICE_ENGINE_DOWN, SYSERR_VS_RPC_FAILURE, SYSERR_VS_NOT_BOUND, SYSERR_VS_DISABLED, SYSERR_VS_INTERNAL_ERROR, SYSERR_VS_SCALEOUT_ERROR, SYSERR_VS_SCALEIN_ERROR, SYSERR_VS_MIGRATE_ERROR, SYSERR_VS_MIGRATE_SCALEOUT_ERROR, SYSERR_VS_MIGRATE_SCALEIN_ERROR, SYSERR_VS_AWAIT_STATIC_SE, SYSERR_VS_MIN_SE_NOT_ASSIGNED, SYSERR_VS_SE_NOT_AT_CURRENT_VERSION, SYSERR_VS_RUNTIME_ABSENT, SYSERR_VS_STATEDB_ERR, SYSERR_VS_SNI_CHILD_PARENT_SELIST_MISMATCH, SYSERR_VS_SNI_PARENT_NOT_FOUND, SYSERR_VS_SNI_CHILD_PARENT_SEGROUP_MISMATCH, SYSERR_SE_MGR_VNIC_ALLOC_FAIL, SYSERR_SE_MGR_VNIC_NOT_FOUND, SYSERR_SE_MGR_UNKNOWN_SE, SYSERR_SE_MGR_UNKNOWN_STATE_TRANSITION, SYSERR_SE_MGR_SE_OFFLINE_HB_FAILURE, SYSERR_SE_UPGRADE_IN_PROGRESS, SYSERR_SE_NOT_CONNECTED, SYSERR_RM_RES_UNAVAIL, SYSERR_RM_RES_UNAVAIL_NOTIFY, SYSERR_RM_RES_NOT_INUSE, SYSERR_RM_CONSUMER_NOT_FOUND, SYSERR_RM_REACHABILITY_FAILED, SYSERR_RM_RELEASE_SE_UNAVAIL, SYSERR_RM_UNKNOWN_SE_GROUP, SYSERR_RM_NO_SE_FOUND, SYSERR_RM_PARTIAL_SE_FOUND, SYSERR_RM_AWAIT_VM_CREATE, SYSERR_RM_AWAIT_VNIC_ADD, SYSERR_RM_AWAIT_BOOTUP, SYSERR_RM_RESOURCE_NOT_FOUND, SYSERR_RM_CANNOT_SPAWN_SE, SYSERR_RM_RES_NOT_NEEDED, SYSERR_RM_RES_INFRA_DELETED, SYSERR_RM_RES_USER_DELETED, SYSERR_RM_RES_USER_REBOOTED, SYSERR_RM_RES_CRASHED, SYSERR_RM_RES_CONN_LOST, SYSERR_RM_RES_VIP_REACH_LOST, SYSERR_RM_VS_PROCESSING, SYSERR_RM_VNIC_IP_FAILURE, SYSERR_RM_STATIC_NO_POOL, SYSERR_RM_STATIC_POOL_EXHAUSTED, SYSERR_RM_VIP_MULT_NETWORKS, SYSERR_RM_SRVR_MULT_NETWORKS, SYSERR_RM_VIP_NO_NETWORK, SYSERR_RM_SRVR_NO_NETWORK, SYSERR_RM_MAX_PARALLEL_SE_CREATE, SYSERR_RM_MAX_SE_CREATE_ATTEMPTS, SYSERR_RM_MULT_SE_CRASH, SYSERR_RM_VS_SE_CREATE_IN_PROG, SYSERR_RM_VS_SE_BOOTUP_IN_PROG, SYSERR_RM_VS_SE_VNIC_ADD_IN_PROG, SYSERR_RM_VS_SE_VNIC_IP_IN_PROG, SYSERR_RM_NO_SUITABLE_HOST, SYSERR_RM_NO_SE_IN_SE_GRP, SYSERR_RM_ALL_SE_IN_SE_GRP_DOWN, SYSERR_RM_NO_SE_IN_SE_GRP_SRVR_ACC, SYSERR_RM_NO_SE_IN_SE_GRP_VIP_ACC, SYSERR_RM_ALL_SE_IN_SE_GRP_MAX_VS, SYSERR_RM_ALL_SE_IN_SE_GRP_NW_ACC_MAX_VS, SYSERR_RM_VIP_SE_NW_ACC, SYSERR_RM_VIP_SE_MAX_VS, SYSERR_RM_VIP_SE_GRP_MISMATCH, SYSERR_RM_VIP_SE_PENDING_OP, SYSERR_RM_MULT_MGMT_SUBNET, SYSERR_RM_MAX_SE_IN_GRP, SYSERR_RM_BOOTUP_FAILURE, SYSERR_RM_PENDING_VNIC_OP, SYSERR_RM_SE_MGMT_NO_STATIC_IPS_CONFIGURED, SYSERR_RM_SE_MGMT_STATIC_IPS_EXHAUSTED, SYSERR_RM_NO_MGMT_SUBNET, SYSERR_RM_MGMT_DHCP_FAILURE, SYSERR_RM_CANNOT_ADD_VNICS, SYSERR_RM_CONSUMER_RESOURCES_SATISFIED, SYSERR_RM_DATA_DHCP_FAILURE, SYSERR_RM_QUERY_HOST_IN_PROGRESS, SYSERR_RM_INSUFFICIENT_BUFFER_SE, SYSERR_RM_NO_DEFAULT_GW_SE_MGMT_NW, SYSERR_RM_PARENT_SE_NW_ACC, SYSERR_RM_PARENT_SE_MAX_VS, SYSERR_RM_PARENT_SE_GRP_MISMATCH, SYSERR_RM_DEF_GW_INCORRECT, SYSERR_RM_NETWORK_NOT_FOUND, SYSERR_RM_ALL_SE_IN_SE_GRP_USED, SYSERR_RM_SE_GRP_PENDING_OP, SYSERR_RM_ALL_SE_IN_SE_GRP_DISABLED, SYSERR_RM_VS_SE_PING_CHECK_IN_PROG, SYSERR_RM_CONSUMER_PENDING_TASK, SYSERR_RM_SE_GRP_VIP_NW_ACC, SYSERR_RM_SE_GRP_NW_ACC, SYSERR_RM_SE_GRP_MAX_VS, SYSERR_RM_ALL_SE_IN_SE_GRP_GW_DOWN, SYSERR_RM_SE_GW_DOWN, SYSERR_RM_SE_DISCONNECTED, SYSERR_RM_RES_USER_DISABLED_FORCE, SYSERR_RM_VS_SE_ATTACH_IP_IN_PROG, SYSERR_RM_LICENSE_EXCEEDED_CANNOT_SPAWN_SE, SYSERR_RM_RES_SWTICHOVER_FORCE, SYSERR_VI_MGR_SEVM_VNIC_SUCCESS, SYSERR_VI_MGR_SEVM_CREATE_FAIL_NO_HW_INFO, SYSERR_VI_MGR_SEVM_CREATE_FAIL_DUPLICATE_NAME, SYSERR_VI_MGR_SEVM_CREATE_FAIL_NO_MGMT_NW, SYSERR_VI_MGR_SEVM_CREATE_FAIL_NO_CPU, SYSERR_VI_MGR_SEVM_CREATE_FAIL_NO_MEM, SYSERR_VI_MGR_SEVM_CREATE_FAIL_NO_LEASE, SYSERR_VI_MGR_SEVM_CREATE_FAIL_OVF_ERROR, SYSERR_VI_MGR_SEVM_CREATE_NO_HOST_VM_NETWORK, SYSERR_VI_MGR_SEVM_CREATE_FAIL_NO_PROGRESS, SYSERR_VI_MGR_SEVM_CREATE_FAIL_ABORTED, SYSERR_VI_MGR_SEVM_CREATE_FAILURE, SYSERR_VI_MGR_SEVM_CREATE_FAIL_POWER_ON, SYSERR_VI_MGR_SEVM_VNIC_NO_VM, SYSERR_VI_MGR_SEVM_VNIC_MAC_ADDR_ERROR, SYSERR_VI_MGR_SEVM_VNIC_FAILURE, SYSERR_VI_MGR_SEVM_VNIC_NO_PG_PORTS, SYSERR_VI_MGR_SEVM_DELETE_FAILURE, SYSERR_VI_MGR_SEVM_CREATE_LIMIT_REACHED, SYSERR_VI_MGR_SEVM_SET_MGMT_IP_FAILED, SYSERR_VI_MGR_SEVM_CREATE_ACCESS_ERROR, SYSERR_VI_MGR_SEVM_CREATE_NO_IMAGE, SYSERR_VI_MGR_SEVM_VINFRA_UNINITIALIZED, SYSERR_VI_MGR_SEVM_CREATE_NO_HOST, SYSERR_VI_MGR_SEVM_CREATE_FAIL_NO_MGMT_NW_PORTS, SYSERR_VI_MGR_SEVM_INVALID_DATA, SYSERR_VI_MGR_SEVM_CREATE_FAIL_MULTIPLE_MGMT_NW, SYSERR_VI_MGR_SEVM_VCENTER_CONN_FAIL, SYSERR_VI_MGR_SEVM_TIMED_OUT, SYSERR_VI_MGR_SEVM_NO_SOURCE_CLONE, SYSERR_VI_MGR_SEVM_NO_AVAILABILITY_ZONE, SYSERR_VI_MGR_SEVM_FLAVOR_UNAVAIL, SYSERR_VI_MGR_SEVM_DELETED, SYSERR_VI_MGR_SEVM_VINFRA_FAILURE, SYSERR_VI_MGR_SEVM_VNIC_FAILURE_QUESTION, SYSERR_VI_MGR_LOGIN_FAIL_NO_VCENTER, SYSERR_VI_MGR_LOGIN_FAIL_USER_CREDENTIALS, SYSERR_VI_MGR_VCENTER_VERSION_MISMATCH, SYSERR_DB_CACHE_TBL_NOT_FOUND, SYSERR_DB_CACHE_OBJ_NOT_FOUND, SYSERR_DB_QUERY_QUEUED, SYSERR_DB_QUERY_BATCHED, SYSERR_DB_UPDATE_FAILED, SYSERR_DB_QUERY_FAILED, SYSERR_OS_AGENT_Q_FULL, SYSERR_OS_AGENT_OPENSTACK_UNINITIALIZED, SYSERR_OS_AGENT_OPENSTACK_ACCESSERR, SYSERR_OS_AGENT_OPENSTACK_RESOURCEERR, SYSERR_OS_AGENT_TENANT_ABSENT, SYSERR_OS_AGENT_INVALID_DATA, SYSERR_CC_SVC_Q_FULL, SYSERR_CC_AGENT_UNINITIALIZED, SYSERR_CC_AGENT_ACCESSERR, SYSERR_CC_AGENT_RESOURCEERR, SYSERR_CC_AGENT_TENANT_ACCESSERR, SYSERR_CC_AGENT_TENANT_ABSENT, SYSERR_CC_SVC_INVALID_DATA, SYSERR_CC_OS_AGENT_NEUTRON_HOST_ACCESSERR, SYSERR_CC_NO_FLAVOR, SYSERR_CC_AGENT_ABSENT, SYSERR_CC_AGENT_CONFIG_FAILURE, SYSERR_CC_AGENT_DECONFIG_FAILURE, SYSERR_CC_AGENT_NON_INFRA_SEVM, SYSERR_MESOS_DISCOVERY_DEPLOYMENT_FAIL, SYSERR_MESOS_DISCOVERY_TIMEOUT, SYSERR_MARATHON_APP_TERMINATED, SYSERR_MARATHON_INACCESSIBLE, SYSERR_FLEET_API_ERROR, SYSERR_MESOS_SSH_CMD_TIMEOUT, SYSERR_MESOS_SSH_ABORTED, SYSERR_MESOS_SSH_FAILURE, SYSERR_MESOS_SSH_NOTFOUND, SYSERR_CC_AGENT_VNIC_NO_IPS_AVAILABLE, SYSERR_CC_AGENT_VNIC_NO_SUBNETWORK, SYSERR_CC_AGENT_VNIC_FAILURE, SYSERR_CC_AGENT_SCALE_IN_FAILED, SYSERR_CC_AGENT_DS_FAILED, SYSERR_CC_AGENT_SCALE_OUT_FAILED, SYSERR_CC_AGENT_NOT_IMPLEMENTED, SYSERR_CC_AGENT_METHOD_NOT_IMPLEMENTED, SYSERR_CC_AGENT_GENERIC_FAILURE, SYSERR_RUM_TOOMANYSAMPLES, SYSERR_METRICS_TOO_MANY_MSG, SYSERR_METRICS_TOO_MANY_MSG_ACROSS_ENTITIES, SYSERR_ANOMALYZER_NOT_ENOUGH_SAMPLES, SYSERR_AUTOSCALE_REASON_INTELLIGENT_AUTOSCALE, SYSERR_AUTOSCALE_REASON_CONFIG_UPDATE, SYSERR_AUTOSCALE_REASON_POOL_STATE_CHANGE, SYSERR_AUTOSCALE_REASON_ALERT, SYSERR_AUTOSCALEIN_FAILED_LIMIT_EXCEEDED, SYSERR_AUTOSCALEOUT_FAILED_LIMIT_EXCEEDED, SYSERR_AUTOSCALE_IGNORED_AS_WITHIN_COOLDOWN, SYSERR_AUTOSCALE_ORCHESTRATION_TIMEOUT, SYSERR_AUTOSCALE_REASON_NOT_ENOUGH_SERVERS, SYSERR_AUTOSCALE_REASON_TOO_MANY_SERVERS, SYSERR_AUTOSCALE_REASON_ORCHESTRATION_FAILED, SYSERR_AUTOSCALE_REASON_MANUAL, SYSERR_AUTOSCALE_POLICY_NOT_FOUND, SYSERR_LICENSE_FIELD_NAME_NOT_SET, SYSERR_LICENSE_FILE_NOT_FOUND, SYSERR_LICENSE_FIELD_VALID_UNTIL_NOT_SET, SYSERR_LICENSE_INVALID_TIERS, SYSERR_LICENSE_FIELD_LICENSE_ID_NOT_PRESENT, SYSERR_LICENSE_INVALID_VERSION, SYSERR_LICENSE_DECRYPTION_FAILED, SYSERR_LICENSE_ENFORCEMENT_KEY_NOT_VALID, SYSERR_SEAGENT_OBJ_INACTIVE, SYSERR_SEAGENT_OBJ_AWAITING_DP_PROGRAMMING, SYSERR_SEAGENT_OBJ_ACTIVE, SYSERR_SEAGENT_OBJ_GRAPHDB_ERROR, SYSERR_SEAGENT_OBJ_DP_ERROR, SYSERR_SEAGENT_OBJ_DISABLED_RULE_POOL, SYSERR_SEAGENT_EASTWEST_VS_SUBNET_ERROR, SYSERR_SEAGENT_OBJ_NOT_FOUND, SYSERR_SEAGENT_VS_NOT_FOUND, SYSERR_SEAGENT_VS_VRF_ERROR, SYSERR_GSLB_INVALID_MTYPE, SYSERR_GSLB_INVALID_SITE_CREDENTIALS, SYSERR_GSLB_OBJECT_NOT_FOUND, SYSERR_GSLB_INVALID_OPS, SYSERR_GSLB_PARTIAL_SUCCESS, SYSERR_GSLB_FQDN_CONFLICT, SYSERR_GSLB_CLEANUP_IN_PROGRESS, SYSERR_GSLB_METHOD_NOP, SYSERR_GSLB_API_NOT_SUPPORTED_FOR_UNFEDERATED_OBJECTS, SYSERR_GSLB_STATEDB_ERR, SYSERR_GSLB_SERVICE_MEMBER_VIPS_NOT_IN_SYNC, SYSERR_GSLB_SERVICE_MEMBER_DISABLED, SYSERR_GSLB_SITE_DISABLED, SYSERR_GSLB_SERVICE_DISABLED, SYSERR_GSLB_HM_PROXY_DOWN, SYSERR_GSLB_DNS_DISABLED, SYSERR_GSLB_SERVICE_NON_AVI_VIP_INFO_UNAVAILABLE, SYSERR_GSLB_SERVICE_DATAPATH_STATUS_UNAVAILABLE, SYSERR_GSLB_SERVICE_MEMBER_SERVICES_NOT_IN_SYNC, SYSERR_GSLB_SERVICE_INCONSISTENT_APPLICATION_PROFILE, SYSERR_GSLB_SERVICE_INVALID_APPLICATION_PROFILE, SYSERR_GSLB_SERVICE_SP_INCONSISTENT_CONFIGURED_SERVERS, SYSERR_GSLB_SERVICE_SP_INCONSISTENT_OPERATIONAL_SERVERS, SYSERR_GSLB_SERVICE_SP_ALL_SERVERS_DOWN, SYSERR_GSLB_SERVICE_SP_SOME_SERVERS_DOWN, SYSERR_GSLB_CONFIGURED_VS_IS_NOT_A_DNS_VS, SYSERR_GSLB_NOT_CONFIGURED, SYSERR_GSLB_INVALID_SENDER, SYSERR_GSLB_INVALID_SENDER_STATE, SYSERR_GSLB_INVALID_RX_ID, SYSERR_GSLB_INVALID_VIEW_ID, SYSERR_GSLB_GROUP_CONFLICT, SYSERR_GSLB_INVALID_MTYPE_AT_FOLLOWER, SYSERR_GSLB_LEADER_NOT_IN_LIST, SYSERR_GSLB_SERVICE_CTRL_STATUS_UNAVAILABLE, SYSERR_GSLB_SITE_FSM_NULL, SYSERR_GSLB_SITE_FSM_DISABLE_IN_PROGRESS, SYSERR_GSLB_SITE_FSM_DISABLED, SYSERR_GSLB_SITE_FSM_JOIN_IN_PROGRESS, SYSERR_GSLB_SITE_FSM_INIT, SYSERR_GSLB_SITE_FSM_UNREACHABLE, SYSERR_GSLB_SITE_FSM_LEAVE_IN_PROGRESS, SYSERR_GSLB_SITE_FSM_MMODE, SYSERR_GSLB_SITE_ACTIVE_TO_PASSIVE_TRANSITION, SYSERR_GSLB_SITE_PASSIVE_TO_ACTIVE_TRANSITION, SYSERR_GSLB_SITE_MAX_RETRIES_DONE, SYSERR_GSLB_TIMEOUT, SYSERR_GSLB_CONNECTION_TIMEOUT, SYSERR_GSLB_CONNECTION_REFUSED_ERROR, SYSERR_GSLB_SERVICE_CTRL_STATUS_NA_DUE_TO_UNREACHABLE_SITE, SYSERR_GSLB_SERVICE_SP_NO_CONFIGURED_SERVERS, SYSERR_GSLB_INVALID_OBJECT, SYSERR_DNS_POLICY_CREATE_FAIL, SYSERR_DNS_POLICY_UPDATE_FAIL, SYSERR_LCM_CORE_NOT_COPIED_DUE_TO_MAX_LIMIT, SYSERR_LCM_CORE_NOT_COPIED_INSUFFICIENT_DISK_SIZE, SYSERR_LCM_SKIP_SIMILAR_CORE, SYSERR_LCM_CORE_NOT_COPIED_DUE_TO_ERRORS.
	ReasonCode *string `json:"reason_code,omitempty"`

	// related objects corresponding to the events.
	RelatedUuids []string `json:"related_uuids,omitempty"`

	// Number of report_timestamp.
	// Required: true
	ReportTimestamp *int64 `json:"report_timestamp"`

	// tenant of EventLog.
	Tenant *string `json:"tenant,omitempty"`

	//  Field introduced in 17.2.1.
	TenantName *string `json:"tenant_name,omitempty"`
}
