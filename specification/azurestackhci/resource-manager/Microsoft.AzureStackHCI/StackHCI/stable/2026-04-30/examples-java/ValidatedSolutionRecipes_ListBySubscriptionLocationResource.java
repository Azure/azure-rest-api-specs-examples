
/**
 * Samples for ValidatedSolutionRecipes ListBySubscriptionLocationResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/ValidatedSolutionRecipes_ListBySubscriptionLocationResource.json
     */
    /**
     * Sample code: ValidatedSolutionRecipes_ListBySubscriptionLocationResource.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void validatedSolutionRecipesListBySubscriptionLocationResource(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.validatedSolutionRecipes().listBySubscriptionLocationResource("westus2",
            com.azure.core.util.Context.NONE);
    }
}
