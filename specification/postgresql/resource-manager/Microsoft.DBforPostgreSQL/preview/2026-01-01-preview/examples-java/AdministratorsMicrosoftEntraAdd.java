
import com.azure.resourcemanager.postgresqlflexibleserver.models.PrincipalType;

/**
 * Samples for AdministratorsMicrosoftEntra CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/AdministratorsMicrosoftEntraAdd.json
     */
    /**
     * Sample code: Add a server administrator associated to a Microsoft Entra principal.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void addAServerAdministratorAssociatedToAMicrosoftEntraPrincipal(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administratorsMicrosoftEntras().define("oooooooo-oooo-oooo-oooo-oooooooooooo")
            .withExistingFlexibleServer("exampleresourcegroup", "exampleserver").withPrincipalType(PrincipalType.USER)
            .withPrincipalName("exampleuser@contoso.com").withTenantId("tttttttt-tttt-tttt-tttt-tttttttttttt").create();
    }
}
