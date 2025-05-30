from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python session_pools_lifecycle_timed_create_or_update.py

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

    response = client.container_apps_session_pools.begin_create_or_update(
        resource_group_name="rg",
        session_pool_name="testsessionpool",
        session_pool_envelope={
            "identity": {"type": "SystemAssigned"},
            "location": "East US",
            "properties": {
                "containerType": "CustomContainer",
                "customContainerTemplate": {
                    "containers": [
                        {
                            "args": ["-c", "while true; do echo hello; sleep 10;done"],
                            "command": ["/bin/sh"],
                            "image": "repo/testcontainer:v4",
                            "name": "testinitcontainer",
                            "resources": {"cpu": 0.25, "memory": "0.5Gi"},
                        }
                    ],
                    "ingress": {"targetPort": 80},
                    "registryCredentials": {
                        "identity": "/subscriptions/7a497526-bb8d-4816-9795-db1418a1f977/resourcegroups/test/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testSP",
                        "server": "test.azurecr.io",
                    },
                },
                "dynamicPoolConfiguration": {
                    "lifecycleConfiguration": {"lifecycleType": "OnContainerExit", "maxAlivePeriodInSeconds": 86400}
                },
                "environmentId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
                "managedIdentitySettings": [{"identity": "system", "lifecycle": "Main"}],
                "poolManagementType": "Dynamic",
                "scaleConfiguration": {"maxConcurrentSessions": 500, "readySessionInstances": 100},
                "sessionNetworkConfiguration": {"status": "EgressEnabled"},
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/SessionPools_LifecycleTimed_CreateOrUpdate.json
if __name__ == "__main__":
    main()
