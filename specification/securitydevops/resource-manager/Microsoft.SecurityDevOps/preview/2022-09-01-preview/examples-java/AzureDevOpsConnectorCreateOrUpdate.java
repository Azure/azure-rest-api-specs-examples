import com.azure.resourcemanager.securitydevops.models.AuthorizationInfo;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsConnectorProperties;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsOrgMetadata;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsProjectMetadata;
import java.util.Arrays;

/** Samples for AzureDevOpsConnector CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsConnectorCreateOrUpdate.json
     */
    /**
     * Sample code: AzureDevOpsConnector_CreateOrUpdate.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsConnectorCreateOrUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager
            .azureDevOpsConnectors()
            .define("testconnector")
            .withRegion("West US")
            .withExistingResourceGroup("westusrg")
            .withProperties(
                new AzureDevOpsConnectorProperties()
                    .withAuthorization(new AuthorizationInfo().withCode("00000000000000000000"))
                    .withOrgs(
                        Arrays
                            .asList(
                                new AzureDevOpsOrgMetadata()
                                    .withName("testOrg")
                                    .withProjects(
                                        Arrays
                                            .asList(
                                                new AzureDevOpsProjectMetadata()
                                                    .withName("testProject")
                                                    .withRepos(Arrays.asList("testRepo")))))))
            .create();
    }
}
