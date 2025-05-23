from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python create_or_update_container_app.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WebSiteManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.container_apps.begin_create_or_update(
        resource_group_name="rg",
        name="testcontainerApp0",
        container_app_envelope={
            "kind": "containerApp",
            "location": "East US",
            "properties": {
                "configuration": {"ingress": {"external": True, "targetPort": 3000}},
                "kubeEnvironmentId": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/kubeEnvironments/demokube",
                "template": {
                    "containers": [{"image": "repo/testcontainerApp0:v1", "name": "testcontainerApp0"}],
                    "dapr": {"appPort": 3000, "enabled": True},
                    "scale": {
                        "maxReplicas": 5,
                        "minReplicas": 1,
                        "rules": [
                            {
                                "custom": {"metadata": {"concurrentRequests": "50"}, "type": "http"},
                                "name": "httpscalingrule",
                            }
                        ],
                    },
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/CreateOrUpdateContainerApp.json
if __name__ == "__main__":
    main()
