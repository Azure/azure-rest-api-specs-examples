from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python session_pools_create_or_update.py

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
                },
                "dynamicPoolConfiguration": {"cooldownPeriodInSeconds": 600, "executionType": "Timed"},
                "environmentId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
                "poolManagementType": "Dynamic",
                "scaleConfiguration": {"maxConcurrentSessions": 500, "readySessionInstances": 100},
                "sessionNetworkConfiguration": {"status": "EgressEnabled"},
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/SessionPools_CreateOrUpdate.json
if __name__ == "__main__":
    main()
