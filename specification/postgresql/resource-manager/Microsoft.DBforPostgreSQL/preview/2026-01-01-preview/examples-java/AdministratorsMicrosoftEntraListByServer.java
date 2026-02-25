
/**
 * Samples for AdministratorsMicrosoftEntra ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/AdministratorsMicrosoftEntraListByServer.json
     */
    /**
     * Sample code: List information about all server administrators associated to Microsoft Entra principals.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listInformationAboutAllServerAdministratorsAssociatedToMicrosoftEntraPrincipals(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administratorsMicrosoftEntras().listByServer("exampleresourcegroup", "exampleserver",
            com.azure.core.util.Context.NONE);
    }
}
