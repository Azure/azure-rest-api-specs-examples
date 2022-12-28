from azure.identity import DefaultAzureCredential
from azure.mgmt.dnsresolver import DnsResolverManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-dnsresolver
# USAGE
    python forwarding_rule_list.py

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

    response = client.forwarding_rules.list(
        resource_group_name="sampleResourceGroup",
        dns_forwarding_ruleset_name="sampleDnsForwardingRuleset",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ForwardingRule_List.json
if __name__ == "__main__":
    main()
