from azure.identity import DefaultAzureCredential

from azure.mgmt.security import SecurityCenter

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-security
# USAGE
    python get_jit_network_access_policies_subscription_example.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityCenter(
        credential=DefaultAzureCredential(),
        subscription_id="20ff7fc3-e762-44dd-bd96-b71116dcdc23",
    )

    response = client.jit_network_access_policies.list()
    for item in response:
        print(item)


# x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesSubscription_example.json
if __name__ == "__main__":
    main()
