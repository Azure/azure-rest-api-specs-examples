
import com.azure.resourcemanager.azurestackhci.models.DeploymentMode;
import com.azure.resourcemanager.azurestackhci.models.DownloadOsJobProperties;
import com.azure.resourcemanager.azurestackhci.models.DownloadOsProfile;
import com.azure.resourcemanager.azurestackhci.models.DownloadRequest;
import com.azure.resourcemanager.azurestackhci.models.ProvisioningOsType;

/**
 * Samples for EdgeMachineJobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/EdgeMachineJobs_CreateOrUpdate_DownloadOs.json
     */
    /**
     * Sample code: EdgeMachineJobs_CreateOrUpdate_DownloadOs.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        edgeMachineJobsCreateOrUpdateDownloadOs(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.edgeMachineJobs().define("DownloadOs").withExistingEdgeMachine("ArcInstance-rg", "machine1")
            .withProperties(new DownloadOsJobProperties().withDeploymentMode(DeploymentMode.DEPLOY)
                .withDownloadRequest(new DownloadRequest().withTarget(ProvisioningOsType.AZURE_LINUX).withOsProfile(
                    new DownloadOsProfile().withOsName("AzureLinux").withOsType("AzureLinux").withOsVersion("3.0")
                        .withOsImageLocation("https://aka.ms/aep/azlinux3.0").withVsrVersion("1.0.0")
                        .withImageHash("sha256:a8b1c2d3e4f5a6b7c8d9e0f1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1")
                        .withGpgPubKey("fakeTokenPlaceholder"))))
            .create();
    }
}
