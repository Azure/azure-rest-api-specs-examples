
import com.azure.resourcemanager.datafactory.models.FactoryRepoUpdate;
import com.azure.resourcemanager.datafactory.models.FactoryVstsConfiguration;

/**
 * Samples for Factories ConfigureFactoryRepo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Factories_ConfigureFactoryRepo.json
     */
    /**
     * Sample code: Factories_ConfigureFactoryRepo.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesConfigureFactoryRepo(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.factories().configureFactoryRepoWithResponse("East US", new FactoryRepoUpdate().withFactoryResourceId(
            "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName")
            .withRepoConfiguration(new FactoryVstsConfiguration().withAccountName("ADF").withRepositoryName("repo")
                .withCollaborationBranch("master").withRootFolder("/").withLastCommitId("").withProjectName("project")
                .withTenantId("")),
            com.azure.core.util.Context.NONE);
    }
}
