from azure.identity import DefaultAzureCredential

from azure.mgmt.applicationinsights import ApplicationInsightsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-applicationinsights
# USAGE
    python feature_capabilities_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApplicationInsightsManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.component_feature_capabilities.get(
        resource_group_name="my-resource-group",
        resource_name="my-component",
    )
    print(response)


# x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FeatureCapabilitiesGet.json
if __name__ == "__main__":
    main()
