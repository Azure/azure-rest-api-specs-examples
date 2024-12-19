from azure.identity import DefaultAzureCredential

from azure.mgmt.cognitiveservices import CognitiveServicesManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cognitiveservices
# USAGE
    python get_shared_commitment_plan.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CognitiveServicesManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
    )

    response = client.commitment_plans.get_plan(
        resource_group_name="resourceGroupName",
        commitment_plan_name="commitmentPlanName",
    )
    print(response)


# x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/GetSharedCommitmentPlan.json
if __name__ == "__main__":
    main()
