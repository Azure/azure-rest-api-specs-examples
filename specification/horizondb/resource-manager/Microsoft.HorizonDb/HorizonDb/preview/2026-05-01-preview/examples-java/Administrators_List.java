
/**
 * Samples for HorizonDbAdministrators List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Administrators_List.json
     */
    /**
     * Sample code: List HorizonDB administrators in a cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBAdministratorsInACluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbAdministrators().list("exampleresourcegroup", "examplecluster",
            com.azure.core.util.Context.NONE);
    }
}
