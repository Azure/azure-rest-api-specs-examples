from azure.identity import DefaultAzureCredential
from azure.mgmt.azurearcdata import AzureArcDataManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurearcdata
# USAGE
    python create_or_update_postgres_instance.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureArcDataManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.postgres_instances.begin_create(
        resource_group_name="testrg",
        postgres_instance_name="testpostgresInstance",
        resource={
            "extendedLocation": {
                "name": "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/arclocation",
                "type": "CustomLocation",
            },
            "location": "eastus",
            "properties": {
                "admin": "admin",
                "basicLoginInformation": {"password": "********", "username": "username"},
                "dataControllerId": "dataControllerId",
                "k8sRaw": {
                    "apiVersion": "apiVersion",
                    "kind": "postgresql-12",
                    "metadata": {
                        "creationTimestamp": "2020-08-25T14:55:10Z",
                        "generation": 1,
                        "name": "pg1",
                        "namespace": "test",
                        "resourceVersion": "527780",
                        "selfLink": "/apis/arcdata.microsoft.com/v1alpha1/namespaces/test/postgresql-12s/pg1",
                        "uid": "1111aaaa-ffff-ffff-ffff-99999aaaaaaa",
                    },
                    "spec": {
                        "backups": {
                            "deltaMinutes": 3,
                            "fullMinutes": 10,
                            "tiers": [
                                {
                                    "retention": {"maximums": ["6", "512MB"], "minimums": ["3"]},
                                    "storage": {"volumeSize": "1Gi"},
                                }
                            ],
                        },
                        "engine": {"extensions": [{"name": "citus"}]},
                        "scale": {"shards": 3},
                        "scheduling": {"default": {"resources": {"requests": {"memory": "256Mi"}}}},
                        "service": {"type": "NodePort"},
                        "storage": {
                            "data": {"className": "local-storage", "size": "5Gi"},
                            "logs": {"className": "local-storage", "size": "5Gi"},
                        },
                    },
                    "status": {"externalEndpoint": None, "readyPods": "4/4", "state": "Ready"},
                },
            },
            "sku": {"dev": True, "name": "default", "tier": "Hyperscale"},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/azurearcdata/resource-manager/Microsoft.AzureArcData/preview/2022-03-01-preview/examples/CreateOrUpdatePostgresInstance.json
if __name__ == "__main__":
    main()
