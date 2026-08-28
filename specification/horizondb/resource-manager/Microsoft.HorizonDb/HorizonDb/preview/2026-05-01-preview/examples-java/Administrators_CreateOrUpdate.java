
import com.azure.resourcemanager.horizondb.models.HorizonDbAdministratorPropertiesForAdd;
import com.azure.resourcemanager.horizondb.models.PrincipalTypes;

/**
 * Samples for HorizonDbAdministrators CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Administrators_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a HorizonDB administrator.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        createOrUpdateAHorizonDBAdministrator(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbAdministrators().define("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
            .withExistingCluster("exampleresourcegroup", "examplecluster")
            .withProperties(new HorizonDbAdministratorPropertiesForAdd().withPrincipalName("admin@contoso.com")
                .withPrincipalType(PrincipalTypes.USER).withTenantId("11111111-2222-3333-4444-555555555555"))
            .create();
    }
}
