from azure.identity import DefaultAzureCredential

from azure.mgmt.domainservices import DomainServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-domainservices
# USAGE
    python unsuspend_domain_service.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DomainServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.domain_services.unsuspend(
        resource_group_name="example-rg",
        domain_service_name="example-domainservice",
    )
    print(response)


# x-ms-original-file: 2025-10-01-preview/UnsuspendDomainService.json
if __name__ == "__main__":
    main()
