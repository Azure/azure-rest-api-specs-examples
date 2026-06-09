
/**
 * Samples for MajorVersionUpgradePrecheck Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/MajorVersionUpgradePrecheckGet.json
     */
    /**
     * Sample code: Get information about a major version upgrade precheck validation.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAMajorVersionUpgradePrecheckValidation(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.majorVersionUpgradePrechecks().getWithResponse("exampleresourcegroup", "exampleserver",
            "pppppppp-pppp-pppp-pppp-pppppppppppp", com.azure.core.util.Context.NONE);
    }
}
