
import com.azure.resourcemanager.horizondb.models.MicrosoftEntraAdministratorPropertiesForAdd;
import com.azure.resourcemanager.horizondb.models.PrincipalTypes;

/**
 * Samples for MicrosoftEntraAdministrators CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAdministrators_CreateOrReplace.json
     */
    /**
     * Sample code: Create or replace a Microsoft Entra administrator for a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createOrReplaceAMicrosoftEntraAdministratorForAHorizonDBCluster(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAdministrators().define("11111111-1111-1111-1111-111111111111")
            .withExistingCluster("example-resource-group", "example-cluster")
            .withProperties(new MicrosoftEntraAdministratorPropertiesForAdd()
                .withPrincipalName("database.administrators@contoso.com").withPrincipalType(PrincipalTypes.GROUP)
                .withTenantId("22222222-2222-2222-2222-222222222222"))
            .create();
    }
}
