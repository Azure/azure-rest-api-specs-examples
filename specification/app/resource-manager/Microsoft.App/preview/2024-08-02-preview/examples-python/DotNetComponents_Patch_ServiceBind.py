from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python dot_net_components_patch_service_bind.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerAppsAPIClient(
        credential=DefaultAzureCredential(),
        subscription_id="8efdecc5-919e-44eb-b179-915dca89ebf9",
    )

    response = client.dot_net_components.begin_update(
        resource_group_name="examplerg",
        environment_name="myenvironment",
        name="mydotnetcomponent",
        dot_net_component_envelope={
            "properties": {
                "componentType": "AspireDashboard",
                "configurations": [{"propertyName": "dashboard-theme", "value": "dark"}],
                "serviceBinds": [
                    {
                        "name": "yellowcat",
                        "serviceId": "/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/managedEnvironments/myenvironment/dotNetComponents/yellowcat",
                    }
                ],
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/DotNetComponents_Patch_ServiceBind.json
if __name__ == "__main__":
    main()
