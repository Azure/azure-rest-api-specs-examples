from azure.identity import DefaultAzureCredential
from azure.mgmt.operationsmanagement import OperationsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-operationsmanagement
# USAGE
    python solution_update.py

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

    response = client.solutions.begin_update(
        resource_group_name="rg1",
        solution_name="solution1",
        parameters={"tags": {"Dept": "IT", "Environment": "Test"}},
    ).result()
    print(response)


# x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionUpdate.json
if __name__ == "__main__":
    main()
