import com.azure.resourcemanager.azurestack.models.DeploymentLicenseRequest;

/** Samples for DeploymentLicense Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/DeploymentLicense/Post.json
     */
    /**
     * Sample code: Creates a license that can be used to deploy an Azure Stack device.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void createsALicenseThatCanBeUsedToDeployAnAzureStackDevice(
        com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .deploymentLicenses()
            .createWithResponse(
                new DeploymentLicenseRequest().withVerificationVersion("1"), com.azure.core.util.Context.NONE);
    }
}
