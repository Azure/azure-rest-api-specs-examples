from azure.identity import DefaultAzureCredential
from azure.mgmt.automanage import AutomanageClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automanage
# USAGE
    python delete_configuration_profile.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomanageClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.configuration_profiles.delete(
        resource_group_name="rg",
        configuration_profile_name="customConfigurationProfile",
    )
    print(response)


# x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfile.json
if __name__ == "__main__":
    main()
