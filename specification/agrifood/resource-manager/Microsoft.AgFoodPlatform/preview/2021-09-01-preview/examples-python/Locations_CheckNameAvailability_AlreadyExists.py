from azure.identity import DefaultAzureCredential
from azure.mgmt.agrifood import AgriFoodMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-agrifood
# USAGE
    python locations_check_name_availability_already_exists.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AgriFoodMgmtClient(
        credential=DefaultAzureCredential(),
        solution_id="SOLUTION_ID",
        subscription_id="11111111-2222-3333-4444-555555555555",
    )

    response = client.locations.check_name_availability(
        body={"name": "existingaccountname", "type": "Microsoft.AgFoodPlatform/farmBeats"},
    )
    print(response)


# x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/Locations_CheckNameAvailability_AlreadyExists.json
if __name__ == "__main__":
    main()
