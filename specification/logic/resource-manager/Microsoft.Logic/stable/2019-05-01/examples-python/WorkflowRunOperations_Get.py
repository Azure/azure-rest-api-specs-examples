from azure.identity import DefaultAzureCredential
from azure.mgmt.logic import LogicManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-logic
# USAGE
    python get_a_run_operation.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = LogicManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.workflow_run_operations.get(
        resource_group_name="testResourceGroup",
        workflow_name="testFlow",
        run_name="08586774142730039209110422528",
        operation_id="ebdcbbde-c4db-43ec-987c-fd0f7726f43b",
    )
    print(response)


# x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunOperations_Get.json
if __name__ == "__main__":
    main()
