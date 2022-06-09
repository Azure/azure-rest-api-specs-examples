```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.QuotaBaseProperties;
import com.azure.resourcemanager.machinelearning.models.QuotaUnit;
import com.azure.resourcemanager.machinelearning.models.QuotaUpdateParameters;
import java.util.Arrays;

/** Samples for Quotas Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Quota/update.json
     */
    /**
     * Sample code: update quotas.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void updateQuotas(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .quotas()
            .updateWithResponse(
                "eastus",
                new QuotaUpdateParameters()
                    .withValue(
                        Arrays
                            .asList(
                                new QuotaBaseProperties()
                                    .withId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace1/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs")
                                    .withType("Microsoft.MachineLearningServices/workspaces/quotas")
                                    .withLimit(100L)
                                    .withUnit(QuotaUnit.COUNT),
                                new QuotaBaseProperties()
                                    .withId(
                                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.MachineLearningServices/workspaces/demo_workspace2/quotas/Standard_DSv2_Family_Cluster_Dedicated_vCPUs")
                                    .withType("Microsoft.MachineLearningServices/workspaces/quotas")
                                    .withLimit(200L)
                                    .withUnit(QuotaUnit.COUNT))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.
