from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python dapr_components_delete.py

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

    client.dapr_components.delete(
        resource_group_name="examplerg",
        environment_name="myenvironment",
        component_name="reddog",
    )


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/DaprComponents_Delete.json
if __name__ == "__main__":
    main()
