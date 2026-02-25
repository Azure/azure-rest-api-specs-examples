
/**
 * Samples for AdministratorsMicrosoftEntra Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/AdministratorsMicrosoftEntraGet.json
     */
    /**
     * Sample code: Get information about a server administrator associated to a Microsoft Entra principal.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAServerAdministratorAssociatedToAMicrosoftEntraPrincipal(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.administratorsMicrosoftEntras().getWithResponse("exampleresourcegroup", "exampleserver",
            "oooooooo-oooo-oooo-oooo-oooooooooooo", com.azure.core.util.Context.NONE);
    }
}
