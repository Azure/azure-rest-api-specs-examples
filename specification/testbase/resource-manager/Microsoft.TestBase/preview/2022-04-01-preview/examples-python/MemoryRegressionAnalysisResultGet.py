from azure.identity import DefaultAzureCredential
from azure.mgmt.testbase import TestBase

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-testbase
# USAGE
    python memory_regression_analysis_result_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TestBase(
        credential=DefaultAzureCredential(),
        subscription_id="subscription-id",
    )

    response = client.analysis_results.get(
        resource_group_name="contoso-rg1",
        test_base_account_name="contoso-testBaseAccount1",
        package_name="contoso-package2",
        test_result_name="Windows-10-1909-Test-Id",
        analysis_result_name="memoryRegression",
    )
    print(response)


# x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2022-04-01-preview/examples/MemoryRegressionAnalysisResultGet.json
if __name__ == "__main__":
    main()
