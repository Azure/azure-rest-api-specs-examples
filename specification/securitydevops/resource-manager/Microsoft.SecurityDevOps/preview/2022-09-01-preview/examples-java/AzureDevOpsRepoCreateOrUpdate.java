import com.azure.resourcemanager.securitydevops.models.ActionableRemediation;
import com.azure.resourcemanager.securitydevops.models.ActionableRemediationState;
import com.azure.resourcemanager.securitydevops.models.AzureDevOpsRepoProperties;
import com.azure.resourcemanager.securitydevops.models.RuleCategory;
import com.azure.resourcemanager.securitydevops.models.TargetBranchConfiguration;
import java.util.Arrays;

/** Samples for AzureDevOpsRepo CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/AzureDevOpsRepoCreateOrUpdate.json
     */
    /**
     * Sample code: AzureDevOpsRepo_CreateOrUpdate.
     *
     * @param manager Entry point to SecurityDevOpsManager.
     */
    public static void azureDevOpsRepoCreateOrUpdate(
        com.azure.resourcemanager.securitydevops.SecurityDevOpsManager manager) {
        manager
            .azureDevOpsRepoes()
            .define("myRepo")
            .withExistingProject("westusrg", "testconnector", "myOrg", "myProject")
            .withProperties(
                new AzureDevOpsRepoProperties()
                    .withRepoId("00000000-0000-0000-0000-000000000000")
                    .withRepoUrl("https://dev.azure.com/myOrg/myProject/_git/myRepo")
                    .withActionableRemediation(
                        new ActionableRemediation()
                            .withState(ActionableRemediationState.ENABLED)
                            .withSeverityLevels(Arrays.asList("High"))
                            .withCategories(Arrays.asList(RuleCategory.SECRETS))
                            .withBranchConfiguration(new TargetBranchConfiguration().withNames(Arrays.asList("main")))))
            .create();
    }
}
