from azure.identity import DefaultAzureCredential

from azure.mgmt.dnsresolver import DnsResolverManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dnsresolver
# USAGE
    python dns_resolver_domain_list_bulk_domains_put.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DnsResolverManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="abdd4249-9f34-4cc6-8e42-c2e32110603e",
    )

    response = client.dns_resolver_domain_lists.begin_create_or_update(
        resource_group_name="sampleResourceGroup",
        dns_resolver_domain_list_name="sampleDnsResolverDomainList",
        parameters={"location": "westus2", "properties": {}, "tags": {"key1": "value1"}},
    ).result()
    print(response)


# x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolverDomainList_BulkDomains_Put.json
if __name__ == "__main__":
    main()
