from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.features import FeatureClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource-features
# USAGE
    python register_feature.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = FeatureClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.features.register(
        resource_provider_namespace="Resource Provider Namespace",
        feature_name="feature",
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/features/stable/2021-07-01/examples/registerFeature.json
if __name__ == "__main__":
    main()
