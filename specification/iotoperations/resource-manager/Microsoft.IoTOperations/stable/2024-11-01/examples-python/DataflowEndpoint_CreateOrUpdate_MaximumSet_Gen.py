from azure.identity import DefaultAzureCredential

from azure.mgmt.iotoperations import IoTOperationsMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-iotoperations
# USAGE
    python dataflow_endpoint_create_or_update_maximum_set_gen.py

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

    response = client.dataflow_endpoint.begin_create_or_update(
        resource_group_name="rgiotoperations",
        instance_name="resource-name123",
        dataflow_endpoint_name="resource-name123",
        resource={
            "extendedLocation": {"name": "qmbrfwcpwwhggszhrdjv", "type": "CustomLocation"},
            "properties": {
                "dataExplorerSettings": {
                    "authentication": {
                        "method": "SystemAssignedManagedIdentity",
                        "systemAssignedManagedIdentitySettings": {"audience": "psxomrfbhoflycm"},
                        "userAssignedManagedIdentitySettings": {
                            "clientId": "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
                            "scope": "zop",
                            "tenantId": "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
                        },
                    },
                    "batching": {"latencySeconds": 9312, "maxMessages": 9028},
                    "database": "yqcdpjsifm",
                    "host": "<cluster>.<region>.kusto.windows.net",
                },
                "dataLakeStorageSettings": {
                    "authentication": {
                        "accessTokenSettings": {"secretRef": "sevriyphcvnlrnfudqzejecwa"},
                        "method": "SystemAssignedManagedIdentity",
                        "systemAssignedManagedIdentitySettings": {"audience": "psxomrfbhoflycm"},
                        "userAssignedManagedIdentitySettings": {
                            "clientId": "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
                            "scope": "zop",
                            "tenantId": "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
                        },
                    },
                    "batching": {"latencySeconds": 9312, "maxMessages": 9028},
                    "host": "<account>.blob.core.windows.net",
                },
                "endpointType": "DataExplorer",
                "fabricOneLakeSettings": {
                    "authentication": {
                        "method": "SystemAssignedManagedIdentity",
                        "systemAssignedManagedIdentitySettings": {"audience": "psxomrfbhoflycm"},
                        "userAssignedManagedIdentitySettings": {
                            "clientId": "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
                            "scope": "zop",
                            "tenantId": "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
                        },
                    },
                    "batching": {"latencySeconds": 9312, "maxMessages": 9028},
                    "host": "https://<host>.fabric.microsoft.com",
                    "names": {"lakehouseName": "wpeathi", "workspaceName": "nwgmitkbljztgms"},
                    "oneLakePathType": "Files",
                },
                "kafkaSettings": {
                    "authentication": {
                        "method": "SystemAssignedManagedIdentity",
                        "saslSettings": {"saslType": "Plain", "secretRef": "visyxoztqnylvbyokhtmpdkwes"},
                        "systemAssignedManagedIdentitySettings": {"audience": "psxomrfbhoflycm"},
                        "userAssignedManagedIdentitySettings": {
                            "clientId": "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
                            "scope": "zop",
                            "tenantId": "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
                        },
                        "x509CertificateSettings": {"secretRef": "afwizrystfslkfqd"},
                    },
                    "batching": {"latencyMs": 3679, "maxBytes": 8887, "maxMessages": 2174, "mode": "Enabled"},
                    "cloudEventAttributes": "PassThrough",
                    "compression": "None",
                    "consumerGroupId": "ukkzcjiyenhxokat",
                    "copyMqttProperties": "Enabled",
                    "host": "pwcqfiqclcgneolpewnyavoulbip",
                    "kafkaAcks": "Zero",
                    "partitionStrategy": "Default",
                    "tls": {"mode": "Enabled", "trustedCaCertificateConfigMapRef": "tectjjvukvelsreihwadh"},
                },
                "localStorageSettings": {"persistentVolumeClaimRef": "jjwqwvd"},
                "mqttSettings": {
                    "authentication": {
                        "method": "SystemAssignedManagedIdentity",
                        "serviceAccountTokenSettings": {"audience": "ejbklrbxgjaqleoycgpje"},
                        "systemAssignedManagedIdentitySettings": {"audience": "psxomrfbhoflycm"},
                        "userAssignedManagedIdentitySettings": {
                            "clientId": "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
                            "scope": "zop",
                            "tenantId": "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
                        },
                        "x509CertificateSettings": {"secretRef": "afwizrystfslkfqd"},
                    },
                    "clientIdPrefix": "kkljsdxdirfhwxtkavldekeqhv",
                    "cloudEventAttributes": "PassThrough",
                    "host": "nyhnxqnbspstctl",
                    "keepAliveSeconds": 0,
                    "maxInflightMessages": 0,
                    "protocol": "Mqtt",
                    "qos": 1,
                    "retain": "Keep",
                    "sessionExpirySeconds": 0,
                    "tls": {"mode": "Enabled", "trustedCaCertificateConfigMapRef": "tectjjvukvelsreihwadh"},
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: 2024-11-01/DataflowEndpoint_CreateOrUpdate_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
