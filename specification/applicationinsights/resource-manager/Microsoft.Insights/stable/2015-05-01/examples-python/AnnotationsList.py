from azure.identity import DefaultAzureCredential

from azure.mgmt.applicationinsights import ApplicationInsightsManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-applicationinsights
# USAGE
    python annotations_list.py

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

    response = client.annotations.list(
        resource_group_name="my-resource-group",
        resource_name="my-component",
        start="2018-02-05T00%253A30%253A00.000Z",
        end="2018-02-06T00%253A33A00.000Z",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsList.json
if __name__ == "__main__":
    main()
