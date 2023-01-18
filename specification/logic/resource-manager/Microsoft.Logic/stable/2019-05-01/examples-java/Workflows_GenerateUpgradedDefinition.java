import com.azure.resourcemanager.logic.models.GenerateUpgradedDefinitionParameters;

/** Samples for Workflows GenerateUpgradedDefinition. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_GenerateUpgradedDefinition.json
     */
    /**
     * Sample code: Generate an upgraded definition.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void generateAnUpgradedDefinition(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .generateUpgradedDefinitionWithResponse(
                "test-resource-group",
                "test-workflow",
                new GenerateUpgradedDefinitionParameters().withTargetSchemaVersion("2016-06-01"),
                com.azure.core.util.Context.NONE);
    }
}
