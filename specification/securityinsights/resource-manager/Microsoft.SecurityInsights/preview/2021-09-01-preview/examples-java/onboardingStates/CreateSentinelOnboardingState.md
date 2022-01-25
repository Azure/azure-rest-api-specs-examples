Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for SentinelOnboardingStates Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/onboardingStates/CreateSentinelOnboardingState.json
     */
    /**
     * Sample code: Create Sentinel onboarding state.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createSentinelOnboardingState(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .sentinelOnboardingStates()
            .define("default")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withCustomerManagedKey(false)
            .create();
    }
}
```
