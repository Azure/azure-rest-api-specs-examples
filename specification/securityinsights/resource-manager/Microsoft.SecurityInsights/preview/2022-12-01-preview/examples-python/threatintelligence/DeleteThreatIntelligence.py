from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python delete_threat_intelligence.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="bd794837-4d29-4647-9105-6339bfdb4e6a",
    )

    response = client.threat_intelligence_indicator.delete(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        name="d9cd6f0b-96b9-3984-17cd-a779d1e15a93",
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/threatintelligence/DeleteThreatIntelligence.json
if __name__ == "__main__":
    main()
