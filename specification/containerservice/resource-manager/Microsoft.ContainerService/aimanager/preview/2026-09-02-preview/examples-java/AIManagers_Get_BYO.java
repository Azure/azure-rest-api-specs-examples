
/**
 * Samples for AIManagers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-02-preview/AIManagers_Get_BYO.json
     */
    /**
     * Sample code: Gets an AI Manager resource attached to an existing AKS cluster (bring-your-own).
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void getsAnAIManagerResourceAttachedToAnExistingAKSClusterBringYourOwn(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagers().getByResourceGroupWithResponse("rg1", "aimanager1", com.azure.core.util.Context.NONE);
    }
}
