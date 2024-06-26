from azure.identity import DefaultAzureCredential
from azure.mgmt.operationsmanagement import OperationsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-operationsmanagement
# USAGE
    python solution_create.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = OperationsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.solutions.begin_create_or_update(
        resource_group_name="rg1",
        solution_name="solution1",
        parameters={
            "location": "East US",
            "plan": {"name": "name1", "product": "product1", "promotionCode": "promocode1", "publisher": "publisher1"},
            "properties": {
                "containedResources": [
                    "/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource1",
                    "/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource2",
                ],
                "referencedResources": [
                    "/subscriptions/sub2/resourceGroups/rg2/providers/provider1/resources/resource2",
                    "/subscriptions/sub2/resourceGroups/rg2/providers/provider2/resources/resource3",
                ],
                "workspaceResourceId": "/subscriptions/sub2/resourceGroups/rg2/providers/Microsoft.OperationalInsights/workspaces/ws1",
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionCreate.json
if __name__ == "__main__":
    main()
