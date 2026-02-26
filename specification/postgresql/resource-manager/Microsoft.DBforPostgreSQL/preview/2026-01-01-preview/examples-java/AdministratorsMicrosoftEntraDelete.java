
/**
 * Samples for AdministratorsMicrosoftEntra Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/AdministratorsMicrosoftEntraDelete.json
     */
    /**
     * Sample code: Delete a server administrator associated to a Microsoft Entra principal.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void deleteAServerAdministratorAssociatedToAMicrosoftEntraPrincipal(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administratorsMicrosoftEntras().delete("exampleresourcegroup", "exampleserver",
            "oooooooo-oooo-oooo-oooo-oooooooooooo", com.azure.core.util.Context.NONE);
    }
}
