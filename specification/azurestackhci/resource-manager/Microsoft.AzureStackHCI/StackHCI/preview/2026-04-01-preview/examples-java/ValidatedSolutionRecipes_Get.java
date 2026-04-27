
/**
 * Samples for ValidatedSolutionRecipes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ValidatedSolutionRecipes_Get.json
     */
    /**
     * Sample code: ValidatedSolutionRecipes_Get.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        validatedSolutionRecipesGet(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.validatedSolutionRecipes().getWithResponse("westus2", "10.2408.0", com.azure.core.util.Context.NONE);
    }
}
