
from azure.identity import DefaultAzureCredential
from azure.mgmt.cdn import CdnManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cdn
# USAGE
    python afd_custom_domains_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret 
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CdnManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.afd_custom_domains.get(
        resource_group_name="RG",
        profile_name="profile1",
        custom_domain_name="domain1",
    )
    print(response)


# x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_Get.json
if __name__ == "__main__":
    main()
