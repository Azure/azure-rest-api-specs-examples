from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python container_apps_patch.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerAppsAPIClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.container_apps.begin_update(
        resource_group_name="rg",
        container_app_name="testcontainerapp0",
        container_app_envelope={
            "location": "East US",
            "properties": {
                "configuration": {
                    "dapr": {
                        "appPort": 3000,
                        "appProtocol": "http",
                        "enableApiLogging": True,
                        "enabled": True,
                        "httpMaxRequestSize": 10,
                        "httpReadBufferSize": 30,
                        "logLevel": "debug",
                    },
                    "ingress": {
                        "customDomains": [
                            {
                                "bindingType": "SniEnabled",
                                "certificateId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com",
                                "name": "www.my-name.com",
                            },
                            {
                                "bindingType": "SniEnabled",
                                "certificateId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com",
                                "name": "www.my-other-name.com",
                            },
                        ],
                        "external": True,
                        "ipSecurityRestrictions": [
                            {
                                "action": "Allow",
                                "description": "Allowing all IP's within the subnet below to access containerapp",
                                "ipAddressRange": "192.168.1.1/32",
                                "name": "Allow work IP A subnet",
                            },
                            {
                                "action": "Allow",
                                "description": "Allowing all IP's within the subnet below to access containerapp",
                                "ipAddressRange": "192.168.1.1/8",
                                "name": "Allow work IP B subnet",
                            },
                        ],
                        "stickySessions": {"affinity": "sticky"},
                        "targetPort": 3000,
                        "traffic": [{"label": "production", "revisionName": "testcontainerapp0-ab1234", "weight": 100}],
                    },
                    "maxInactiveRevisions": 10,
                    "runtime": {"java": {"enableMetrics": True}},
                    "service": {"type": "redis"},
                },
                "template": {
                    "containers": [
                        {
                            "image": "repo/testcontainerapp0:v1",
                            "name": "testcontainerapp0",
                            "probes": [
                                {
                                    "httpGet": {
                                        "httpHeaders": [{"name": "Custom-Header", "value": "Awesome"}],
                                        "path": "/health",
                                        "port": 8080,
                                    },
                                    "initialDelaySeconds": 3,
                                    "periodSeconds": 3,
                                    "type": "Liveness",
                                }
                            ],
                        }
                    ],
                    "initContainers": [
                        {
                            "image": "repo/testcontainerapp0:v4",
                            "name": "testinitcontainerApp0",
                            "resources": {"cpu": 0.5, "memory": "1Gi"},
                        }
                    ],
                    "scale": {
                        "cooldownPeriod": 350,
                        "maxReplicas": 5,
                        "minReplicas": 1,
                        "pollingInterval": 35,
                        "rules": [
                            {
                                "custom": {"metadata": {"concurrentRequests": "50"}, "type": "http"},
                                "name": "httpscalingrule",
                            }
                        ],
                    },
                    "serviceBinds": [
                        {
                            "name": "service",
                            "serviceId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/containerApps/service",
                        }
                    ],
                },
            },
            "tags": {"tag1": "value1", "tag2": "value2"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ContainerApps_Patch.json
if __name__ == "__main__":
    main()
