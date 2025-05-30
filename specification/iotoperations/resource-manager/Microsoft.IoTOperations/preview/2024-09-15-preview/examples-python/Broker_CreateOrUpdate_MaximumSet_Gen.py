from azure.identity import DefaultAzureCredential

from azure.mgmt.iotoperations import IoTOperationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotoperations
# USAGE
    python broker_create_or_update_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = IoTOperationsMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.broker.begin_create_or_update(
        resource_group_name="rgiotoperations",
        instance_name="resource-name123",
        broker_name="resource-name123",
        resource={
            "extendedLocation": {"name": "qmbrfwcpwwhggszhrdjv", "type": "CustomLocation"},
            "properties": {
                "advanced": {
                    "clients": {
                        "maxKeepAliveSeconds": 3744,
                        "maxMessageExpirySeconds": 3263,
                        "maxPacketSizeBytes": 3029,
                        "maxReceiveMaximum": 2365,
                        "maxSessionExpirySeconds": 3859,
                        "subscriberQueueLimit": {"length": 6, "strategy": "None"},
                    },
                    "encryptInternalTraffic": "Enabled",
                    "internalCerts": {
                        "duration": "bchrc",
                        "privateKey": {"algorithm": "Ec256", "rotationPolicy": "Always"},
                        "renewBefore": "xkafmpgjfifkwwrhkswtopdnne",
                    },
                },
                "cardinality": {
                    "backendChain": {"partitions": 11, "redundancyFactor": 5, "workers": 15},
                    "frontend": {"replicas": 2, "workers": 6},
                },
                "diagnostics": {
                    "logs": {"level": "rnmwokumdmebpmfxxxzvvjfdywotav"},
                    "metrics": {"prometheusPort": 7581},
                    "selfCheck": {"intervalSeconds": 158, "mode": "Enabled", "timeoutSeconds": 14},
                    "traces": {
                        "cacheSizeMegabytes": 28,
                        "mode": "Enabled",
                        "selfTracing": {"intervalSeconds": 22, "mode": "Enabled"},
                        "spanChannelCapacity": 1000,
                    },
                },
                "diskBackedMessageBuffer": {
                    "ephemeralVolumeClaimSpec": {
                        "accessModes": ["nuluhigrbb"],
                        "dataSource": {"apiGroup": "npqapyksvvpkohujx", "kind": "wazgyb", "name": "cwhsgxxcxsyppoefm"},
                        "dataSourceRef": {
                            "apiGroup": "mnfnykznjjsoqpfsgdqioupt",
                            "kind": "odynqzekfzsnawrctaxg",
                            "name": "envszivbbmixbyddzg",
                            "namespace": "etcfzvxqd",
                        },
                        "resources": {"limits": {"key2719": "hmphcrgctu"}, "requests": {"key2909": "txocprnyrsgvhfrg"}},
                        "selector": {
                            "matchExpressions": [
                                {"key": "e", "operator": "In", "values": ["slmpajlywqvuyknipgztsonqyybt"]}
                            ],
                            "matchLabels": {"key6673": "wlngfalznwxnurzpgxomcxhbqefpr"},
                        },
                        "storageClassName": "sseyhrjptkhrqvpdpjmornkqvon",
                        "volumeMode": "rxvpksjuuugqnqzeiprocknbn",
                        "volumeName": "c",
                    },
                    "maxSize": "500M",
                    "persistentVolumeClaimSpec": {
                        "accessModes": ["nuluhigrbb"],
                        "dataSource": {"apiGroup": "npqapyksvvpkohujx", "kind": "wazgyb", "name": "cwhsgxxcxsyppoefm"},
                        "dataSourceRef": {
                            "apiGroup": "mnfnykznjjsoqpfsgdqioupt",
                            "kind": "odynqzekfzsnawrctaxg",
                            "name": "envszivbbmixbyddzg",
                            "namespace": "etcfzvxqd",
                        },
                        "resources": {"limits": {"key2719": "hmphcrgctu"}, "requests": {"key2909": "txocprnyrsgvhfrg"}},
                        "selector": {
                            "matchExpressions": [
                                {"key": "e", "operator": "In", "values": ["slmpajlywqvuyknipgztsonqyybt"]}
                            ],
                            "matchLabels": {"key6673": "wlngfalznwxnurzpgxomcxhbqefpr"},
                        },
                        "storageClassName": "sseyhrjptkhrqvpdpjmornkqvon",
                        "volumeMode": "rxvpksjuuugqnqzeiprocknbn",
                        "volumeName": "c",
                    },
                },
                "generateResourceLimits": {"cpu": "Enabled"},
                "memoryProfile": "Tiny",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-09-15-preview/Broker_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
