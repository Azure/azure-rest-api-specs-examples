from azure.identity import DefaultAzureCredential

from azure.mgmt.recoveryservices import RecoveryServicesClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-recoveryservices
# USAGE
    python get_private_link_resources.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = RecoveryServicesClient(
        credential=DefaultAzureCredential(),
        subscription_id="6c48fa17-39c7-45f1-90ac-47a587128ace",
    )

    response = client.private_link_resources.get(
        resource_group_name="petesting",
        vault_name="pemsi-ecy-rsv2",
        private_link_resource_name="backupResource",
    )
    print(response)


# x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/GetPrivateLinkResources.json
if __name__ == "__main__":
    main()
