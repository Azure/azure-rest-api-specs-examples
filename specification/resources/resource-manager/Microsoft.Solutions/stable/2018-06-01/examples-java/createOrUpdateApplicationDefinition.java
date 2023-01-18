import com.azure.resourcemanager.managedapplications.models.ApplicationLockLevel;
import com.azure.resourcemanager.managedapplications.models.ApplicationProviderAuthorization;
import java.util.Arrays;

/** Samples for ApplicationDefinitions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplicationDefinition.json
     */
    /**
     * Sample code: Create or update managed application definition.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void createOrUpdateManagedApplicationDefinition(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applicationDefinitions()
            .define("myManagedApplicationDef")
            .withRegion("East US 2")
            .withExistingResourceGroup("rg")
            .withLockLevel(ApplicationLockLevel.NONE)
            .withAuthorizations(
                Arrays
                    .asList(
                        new ApplicationProviderAuthorization()
                            .withPrincipalId("validprincipalguid")
                            .withRoleDefinitionId("validroleguid")))
            .withDisplayName("myManagedApplicationDef")
            .withDescription("myManagedApplicationDef description")
            .withPackageFileUri("https://path/to/packagezipfile")
            .create();
    }
}
