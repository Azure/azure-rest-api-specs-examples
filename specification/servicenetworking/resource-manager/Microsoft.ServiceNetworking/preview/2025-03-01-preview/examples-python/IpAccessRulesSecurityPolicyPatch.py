from azure.identity import DefaultAzureCredential

from azure.mgmt.servicenetworking import ServiceNetworkingMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-servicenetworking
# USAGE
    python ip_access_rules_security_policy_patch.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ServiceNetworkingMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.security_policies_interface.update(
        resource_group_name="rg1",
        traffic_controller_name="tc1",
        security_policy_name="sp1",
        properties={"properties": {"ipAccessRulesPolicy": {"rules": []}}},
    )
    print(response)


# x-ms-original-file: 2025-03-01-preview/IpAccessRulesSecurityPolicyPatch.json
if __name__ == "__main__":
    main()
