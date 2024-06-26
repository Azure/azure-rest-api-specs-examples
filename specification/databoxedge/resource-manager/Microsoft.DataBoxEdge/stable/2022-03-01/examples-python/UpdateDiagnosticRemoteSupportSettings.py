from azure.identity import DefaultAzureCredential
from azure.mgmt.databoxedge import DataBoxEdgeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-databoxedge
# USAGE
    python update_diagnostic_remote_support_settings.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DataBoxEdgeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="4385cf00-2d3a-425a-832f-f4285b1c9dce",
    )

    response = client.diagnostic_settings.begin_update_diagnostic_remote_support_settings(
        device_name="testedgedevice",
        resource_group_name="GroupForEdgeAutomation",
        diagnostic_remote_support_settings={
            "properties": {
                "remoteSupportSettingsList": [
                    {
                        "accessLevel": "ReadWrite",
                        "expirationTimeStampInUTC": "2021-07-07T00:00:00+00:00",
                        "remoteApplicationType": "Powershell",
                    }
                ]
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/UpdateDiagnosticRemoteSupportSettings.json
if __name__ == "__main__":
    main()
