from azure.identity import DefaultAzureCredential
from azure.mgmt.springappdiscovery import SpringAppDiscoveryMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-springappdiscovery
# USAGE
    python springbootapps_get_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SpringAppDiscoveryMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="jnetwlorzmxpxmcucorv",
    )

    response = client.springbootapps.get(
        resource_group_name="rgspringbootapps",
        site_name="pdfosfhtemfsaglvwjdyqlyeipucrd",
        springbootapps_name="ofjeesoahqtnovlbuvflyknpbhcpeqqhekntvqxyemuwbcqnuxjgfhsf",
    )
    print(response)


# x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootapps_Get_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
