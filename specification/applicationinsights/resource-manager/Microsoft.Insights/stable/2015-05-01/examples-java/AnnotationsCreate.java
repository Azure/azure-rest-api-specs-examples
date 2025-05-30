
import com.azure.resourcemanager.applicationinsights.fluent.models.AnnotationInner;
import java.time.OffsetDateTime;

/**
 * Samples for Annotations Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/
     * AnnotationsCreate.json
     */
    /**
     * Sample code: AnnotationsCreate.
     * 
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void
        annotationsCreate(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.annotations().createWithResponse("my-resource-group", "my-component",
            new AnnotationInner().withAnnotationName("TestAnnotation").withCategory("Text")
                .withEventTime(OffsetDateTime.parse("2018-01-31T13:41:38.657Z"))
                .withId("444e2c08-274a-4bbb-a89e-d77bb720f44a")
                .withProperties("{\"Comments\":\"Testing\",\"Label\":\"Success\"}"),
            com.azure.core.util.Context.NONE);
    }
}
