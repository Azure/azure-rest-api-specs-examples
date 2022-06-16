import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.applicationinsights.models.WorkbookTemplateGallery;
import java.io.IOException;
import java.util.Arrays;

/** Samples for WorkbookTemplates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateAdd.json
     */
    /**
     * Sample code: WorkbookTemplateAdd.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookTemplateAdd(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) throws IOException {
        manager
            .workbookTemplates()
            .define("testtemplate2")
            .withRegion("west us")
            .withExistingResourceGroup("my-resource-group")
            .withPriority(1)
            .withAuthor("Contoso")
            .withTemplateData(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize(
                        "{\"$schema\":\"https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json\",\"items\":[{\"name\":\"text"
                            + " - 2\",\"type\":1,\"content\":{\"json\":\"## New workbook\\n"
                            + "---\\n"
                            + "\\n"
                            + "Welcome to your new workbook.  This area will display text formatted as markdown.\\n"
                            + "\\n"
                            + "\\n"
                            + "We've included a basic analytics query to get you started. Use the `Edit` button below"
                            + " each section to configure it or add more sections.\"}},{\"name\":\"query -"
                            + " 2\",\"type\":3,\"content\":{\"exportToExcelOptions\":\"visible\",\"query\":\"union"
                            + " withsource=TableName *\\n"
                            + "| summarize Count=count() by TableName\\n"
                            + "| render"
                            + " barchart\",\"queryType\":0,\"resourceType\":\"microsoft.operationalinsights/workspaces\",\"size\":1,\"version\":\"KqlItem/1.0\"}}],\"styleSettings\":{},\"version\":\"Notebook/1.0\"}",
                        Object.class,
                        SerializerEncoding.JSON))
            .withGalleries(
                Arrays
                    .asList(
                        new WorkbookTemplateGallery()
                            .withName("Simple Template")
                            .withCategory("Failures")
                            .withType("tsg")
                            .withOrder(100)
                            .withResourceType("microsoft.insights/components")))
            .create();
    }
}
