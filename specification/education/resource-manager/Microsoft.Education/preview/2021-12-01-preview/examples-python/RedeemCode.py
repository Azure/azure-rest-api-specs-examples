from azure.identity import DefaultAzureCredential
from azure.mgmt.education import EducationManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-education
# USAGE
    python redeem_code.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = EducationManagementClient(
        credential=DefaultAzureCredential(),
    )

    response = client.redeem_invitation_code(
        parameters={"firstName": "test", "lastName": "user", "redeemCode": "exampleRedeemCode"},
    )
    print(response)


# x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/RedeemCode.json
if __name__ == "__main__":
    main()
