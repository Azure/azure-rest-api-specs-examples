from azure.identity import DefaultAzureCredential

from azure.mgmt.web import WebSiteManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-web
# USAGE
    python workflow_run_action_repetitions_request_histories_get.py

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

    response = client.workflow_run_action_repetitions_request_histories.get(
        resource_group_name="test-resource-group",
        name="test-name",
        workflow_name="test-workflow",
        run_name="08586776228332053161046300351",
        action_name="HTTP_Webhook",
        repetition_name="000001",
        request_history_name="08586611142732800686",
    )
    print(response)


# x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/WorkflowRunActionRepetitionsRequestHistories_Get.json
if __name__ == "__main__":
    main()
