from azure.identity import DefaultAzureCredential
from azure.mgmt.labservices import ManagedLabsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-labservices
# USAGE
    python get_lab_plan.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedLabsClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.lab_plans.get(
        resource_group_name="testrg123",
        lab_plan_name="testlabplan",
    )
    print(response)


# x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/LabPlans/getLabPlan.json
if __name__ == "__main__":
    main()
