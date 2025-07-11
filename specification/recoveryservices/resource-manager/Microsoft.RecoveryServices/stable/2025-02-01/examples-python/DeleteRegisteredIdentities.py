from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservices import RecoveryServicesClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservices
# USAGE
    python delete_registered_identities.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesClient(
        credential=DefaultAzureCredential(),
        subscription_id="77777777-d41f-4550-9f70-7708a3a2283b",
    )

    client.registered_identities.delete(
        resource_group_name="BCDRIbzRG",
        vault_name="BCDRIbzVault",
        identity_name="dpmcontainer01",
    )


# x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/DeleteRegisteredIdentities.json
if __name__ == "__main__":
    main()
