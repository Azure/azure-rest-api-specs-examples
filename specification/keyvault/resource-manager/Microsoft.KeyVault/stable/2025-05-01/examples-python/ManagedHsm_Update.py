from azure.identity import DefaultAzureCredential

from azure.mgmt.keyvault import KeyVaultManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-keyvault
# USAGE
    python managed_hsm_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = KeyVaultManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.managed_hsms.begin_update(
        resource_group_name="hsm-group",
        name="hsm1",
        parameters={"tags": {"Dept": "hsm", "Environment": "dogfood", "Slice": "A"}},
    ).result()
    print(response)


# x-ms-original-file: 2025-05-01/ManagedHsm_Update.json
if __name__ == "__main__":
    main()
