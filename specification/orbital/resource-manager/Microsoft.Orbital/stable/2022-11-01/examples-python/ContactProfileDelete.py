from azure.identity import DefaultAzureCredential
from azure.mgmt.orbital import AzureOrbital

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-orbital
# USAGE
    python contact_profile_delete.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureOrbital(
        credential=DefaultAzureCredential(),
        subscription_id="c1be1141-a7c9-4aac-9608-3c2e2f1152c3",
    )

    response = client.contact_profiles.begin_delete(
        resource_group_name="contoso-Rgp",
        contact_profile_name="CONTOSO-CP",
    ).result()
    print(response)


# x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactProfileDelete.json
if __name__ == "__main__":
    main()
